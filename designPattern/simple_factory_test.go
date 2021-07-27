package designPattern

/**
工厂模式
优点：只需要传入一个正确的参数，就可以获取你所需要的对象而无需知道其创建对象的细节

缺点：扩展性差，当增加新的产品需要修改工厂类的判断逻辑，违背开闭原则，如我想要买一个三星手机的话，除了新增三星手机这个产品类，还需要修改工厂中的逻辑
*/

import (
	"fmt"
	"testing"
)

type Phone interface {
	Produce()
}

type ApplePhone struct{}

func (a ApplePhone) Produce() {
	fmt.Println("苹果手机")
}

type HuaWeiPhone struct{}

func (h HuaWeiPhone) Produce() {
	fmt.Println("华为手机")
}

type Factory struct{}

func (f Factory) GetPhone(name string) Phone {
	var phone Phone
	if name == "apple" {
		phone = new(ApplePhone)
	} else if name == "huawei" {
		phone = new(HuaWeiPhone)
	}
	return phone
}

func TestFactory(t *testing.T) {
	f := Factory{}
	applePhone := f.GetPhone("apple")
	applePhone.Produce()
	huaweiPhone := f.GetPhone("huawei")
	huaweiPhone.Produce()
}
