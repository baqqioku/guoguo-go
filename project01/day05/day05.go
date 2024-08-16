package day05

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("call nokia phone")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iPhone ,I can call you")
}

func InterfaceUp() {
	var phone Phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func InterfaceMock() {
	var s Shape = Rectangle{width: 10, height: 10}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle{radius: 5}
	fmt.Print(s.area())
	fmt.Printf("圆形面积: %f\n", s.area())
}
