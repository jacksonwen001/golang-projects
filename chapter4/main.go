package main

import "fmt"
// 定义接口
type CustomerInterface interface {
	sayHi()
}

// 实现接口。 字段定义
type CustomerService struct {

}
// 方法实现
func (c CustomerService) sayHi() {
	fmt.Println("Hello")
}

// 实例化，默认只有一个实例存在
func NewCustomerService() CustomerInterface {
	return &CustomerService{}
}

func main() {
	fmt.Println("Reflection and Interfaces")
	fmt.Println("反射字段，反射并执行方法")
	fmt.Println("struct 定义字段， interface 定义方法。 继承使用鸭子模式进行实现")
	c := NewCustomerService()
	c.sayHi()
}
