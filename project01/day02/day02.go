package day02

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func NoName() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}

func TestNoneName() {
	add := func(a int, b int) int {
		return a + b
	}
	// 调用匿名函数
	result := add(3, 5)

	fmt.Println(result)

	//将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2+8=", sum)

	//也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)

	fmt.Println("10-4=", difference)

}
