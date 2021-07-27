package designPattern

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/rpc"
	"testing"
	"time"
)

/**
代理模式
代理模式为一个对象提供一种代理以控制对该对象的访问
*/

type ProxyRecord struct {
	Key   string
	Value string
}

type KvDb interface {
	// 存储数据
	// 其中reply为操作结果，存储成功为true，否则为false
	// 当连接数据库失败时返回error，成功则返回nil
	Save(record ProxyRecord, reply *bool) error
	// 根据key获取value，其中value通过函数参数中指针类型返回
	// 当连接数据库失败时返回error，成功则返回nil
	Get(key string, value *string) error
}

type Server struct {
	// 采用map存储key-value数据
	data map[string]string
}

func (s *Server) Save(record ProxyRecord, reply *bool) error {
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[record.Key] = record.Value
	*reply = true
	return nil
}

func (s *Server) Get(key string, reply *string) error {
	val, ok := s.data[key]
	if !ok {
		*reply = ""
		return errors.New("DB has no key " + key)
	}
	*reply = val
	return nil
}

//消息处理系统和数据库并不在同一台机器上，因此消息处理系统不能直接调用db.Server的方法进行数据存储，
//像这种服务提供者和服务使用者不在同一机器上的场景，使用远程代理再适合不过了。

func Start()  {
	rpcServer := rpc.NewServer()
	server := &Server{data: make(map[string]string)}
	if err := rpcServer.Register(server);err!=nil{
		fmt.Printf("Register Server to rpc failed, error: %v", err)
		return
	}
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Listen tcp failed, error: %v", err)
		return
	}
	go rpcServer.Accept(l)
	time.Sleep(1000 * time.Second)
	fmt.Println("Rpc server start success.")
}

//到目前为止，我们已经为数据库提供了对外访问的方式。
//现在，我们需要一个远程代理来连接数据库服务端，并进行相关的数据库操作。
//对消息处理系统而言，它不需要，也不应该知道远程代理与数据库服务端交互的底层细节，这样可以减轻系统之间的耦合。
//因此，远程代理需要实现db.KvDb：

type Client struct{
	cli *rpc.Client
}

func (c *Client) Save(record ProxyRecord, reply *bool) error {
	var ret bool
	err := c.cli.Call("Server.Save", record, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Save rpc failed, error: %v", err)
		*reply = false
		return err
	}
	*reply = ret
	return nil
}

func (c *Client) Get(key string, reply *string) error {
	var ret string
	err := c.cli.Call("Server.Get", key, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Get rpc failed, error: %v", err)
		*reply = ""
		return err
	}
	*reply = ret
	return nil
}

func CreateClient() *Client {
	rpcCli, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Create rpc client failed, error: %v.", err)
		return nil
	}
	return &Client{cli: rpcCli}
}

func TestProxy(t *testing.T) {
	go Start()
	cli := CreateClient()
	var sval bool
	err := cli.Save(ProxyRecord{
		Key:   "db",
		Value: "proxy",
	},&sval)
	if err != nil {
		t.Errorf("Save db failed, error: %v\n.", err)
	}
	if !sval {
		t.Errorf("Save db failed, sval: %v\n.", sval)
	}
	var val string
	err = cli.Get("db",&val)
	if err != nil {
		t.Errorf("Get db failed, error: %v\n.", err)
	}

	if val != "proxy" {
		t.Errorf("expect HELLO WORLD, but actual %s.", val)
	}
}