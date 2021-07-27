package designPattern

import (
	"github.com/magiconair/properties/assert"
	"sync"
	"testing"
)

/**
单例模式
保证一个类仅有一个实例，并提供一个访问它的全局访问点
 */
type Message struct {
	Count int
}

type messagePool struct {
	pool *sync.Pool
}

// 饿汉模式：实例在系统加载时候已经完成初始化
var msgPool1 = &messagePool{
	pool: &sync.Pool{
		New: func() interface{} { return &Message{Count: 0} },
	},
}

func Instance1() *messagePool {
	return msgPool1
}

func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

func (m messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}

// ---------------------------------------------------------
// 懒汉模式
var once = sync.Once{}
var msgPool2 *messagePool

func Instance2() *messagePool {
	once.Do(func() {
		msgPool2 = &messagePool{
			pool: &sync.Pool{
				New: func() interface{} {
					return &Message{Count: 0}
				},
			},
		}
	})
	return msgPool2
}

func Test_pool(t *testing.T) {
	// 饿汉模式
	msg := Instance1().GetMsg()
	assert.Equal(t, msg.Count, 0)
	msg.Count = 1
	Instance1().AddMsg(msg)
	msg = Instance1().GetMsg()
	assert.Equal(t, msg.Count, 1)

	// 懒汉模式
	lazyMsg := Instance2().GetMsg()
	assert.Equal(t, lazyMsg.Count, 0)
	lazyMsg.Count = 1
	Instance2().AddMsg(msg)
	msg = Instance2().GetMsg()
	assert.Equal(t, msg.Count, 1)
}
