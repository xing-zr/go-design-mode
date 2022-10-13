package main

import "fmt"

//适配的目标
type V5 interface {
	UseV5()
}

//业务类，依赖V5接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.UseV5()
}

//被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) UseV5() {
	fmt.Println("使用适配器进行充电")
	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
