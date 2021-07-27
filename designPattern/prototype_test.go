package designPattern

/**
原型模式
原型模式主要解决对象复制的问题，它的核心就是clone()方法，返回Prototype对象的复制品。
 */

type Prototype interface {
	clone() Prototype
}

type PrototypeHeader struct {

}
type PrototypeBody struct {

}

type PrototypeMessage struct {
	Header *PrototypeHeader
	Body *PrototypeBody
}

func (p *PrototypeMessage)clone() PrototypeMessage {
	msg := *p
	return msg
}