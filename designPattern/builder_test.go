package designPattern

import (
	"github.com/magiconair/properties/assert"
	"sync"
	"testing"
)

type BuilderHeader struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type BuilderBody struct {
	Items []string
}

type BuilderMessage struct {
	Header *BuilderHeader
	Body   *BuilderBody
}

// 直接创建，对对象使用者不友好和代码可读性差的缺点
func Test_BuilderMessage_d(t *testing.T) {
	message := BuilderMessage{
		Header: &BuilderHeader{
			SrcAddr:  "",
			SrcPort:  0,
			DestAddr: "",
			DestPort: 0,
			Items:    make(map[string]string),
		},
		Body: &BuilderBody{Items: make([]string, 0)},
	}
	// 字段操作
	message.Header.Items["content"] = "json"
	message.Body.Items = append(message.Body.Items, "message")
}

type builder struct {
	once *sync.Once
	msg  *BuilderMessage
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg: &BuilderMessage{
			Header: &BuilderHeader{},
			Body:   &BuilderBody{},
		},
	}
}
func (b *builder) WithSrcAddr(src string) *builder {
	b.msg.Header.SrcAddr = src
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

func (b *builder) WithHeaderItem(key, value string) *builder {
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) Build() *BuilderMessage {
	return b.msg
}

func Test_Builder(t *testing.T) {
	msg := Builder().
		WithSrcAddr("localhost").
		WithHeaderItem("key", "value").
		WithBodyItem("body").
		Build()
	assert.Equal(t, msg.Header.SrcAddr, "localhost")
	assert.Equal(t, msg.Header.Items["key"], "value")
	assert.Equal(t, msg.Body.Items[0], "body")
}
