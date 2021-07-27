package designPattern

import "testing"

/**
组合模式
1. 直接组合
2. 嵌入组合
 */

// 直接组合
type CompositeHeader struct {
	SrcAddr string
}

type CompositeBody struct {}

type CompositeMessage1 struct {
	Header *CompositeHeader
	Body   *CompositeBody
}

// 嵌入组合
type CompositeMessage2 struct {
	CompositeHeader
	CompositeBody
}

func TestComposite(t *testing.T) {
	msg1 := &CompositeMessage1{}
	msg1.Header.SrcAddr = "localhost"

	msg2 := &CompositeMessage2{}
	msg2.SrcAddr = "localhost"
}

// --------------------------------------------

type Car interface {
	Start()
	Stop()
	Fix()
}

type Tractor struct {
	working bool
}
func (t *Tractor) Start() {t.working = true}
func (t *Tractor) Stop()  {t.working = true}
func (t *Tractor) Fix()  {}
func (t *Tractor) Work()  {}

type SportsCar struct {
	working bool
}
func (s *SportsCar) Start() {s.working = true}
func (s *SportsCar) Stop()  {s.working = true}
func (s *SportsCar) Fix()  {}
func (s *SportsCar) DragRacing()  {}

// 每新增一个新的类型，都需要实现这三个方法。但是大多数插件的这三个方法的逻辑基本一致，因此导致了一定程度的代码冗余。
//对于重复代码问题，有什么好的解决方法呢？组合模式！
//下面，我们使用组合模式将这个方法提取成一个新的对象BaseFunc，这样新增一个类型时，只需将BaseFunc作为匿名成员（嵌入组合），就能解决冗余代码问题了。

type BaseFunc struct {
	working bool
}

func (b *BaseFunc) Start() {
	b.working = true
}

func (b *BaseFunc) Stop()  {
	b.working = true
}

func (b *BaseFunc) Fix() {
}

type Truck struct {
	BaseFunc
	transport string
}

func (t *Truck) Transport(goods string)  {
	t.transport = goods
}

func TestCompositeInterface(t *testing.T) {
	var car Car
	truck := &Truck{}
	truck.Transport("water")
	car = truck
	car.Start()
}