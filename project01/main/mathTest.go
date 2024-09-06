package main

import (
	"fmt"
	"project01/day03"
	"project01/day04"
	"project01/day05"
	"project01/day06"
	"time"
)

var (
	guoguo string
	x      int
	y      int
)

const (
	n1 = 100
	n2 = 200
	n3 = 300
	a  = iota
	b
	c
)

func main() {

	//fmt.Println(math.Sin(1))
	//test1()
	//testIota()
	//testIosta2()
	//testLeft()
	//testLoop()
	//testBreakLoop()
	//testBreak2()
	//testSwitch()
	//day02.NoName()
	//day02.TestNoneName()
	//day03.TestArray()
	//day03.Balance()
	day03.VariaArray()
	day03.ManyString()

	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = day03.GetAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)

	var z int = 100
	var x int = 200

	day04.Swap(&z, &x)

	fmt.Printf("交换后，z 的值 : %d\n", z)
	fmt.Printf("交换后，x 的值 : %d\n", x)

	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray)
	day04.ModifyArray(myArray)

	fmt.Println(myArray)
	day04.ModifyArrayWithPointer(&myArray)
	fmt.Println(myArray)

	day04.Pptr()

	day04.MakeSlice()

	day05.InterfaceUp()
	day05.InterfaceMock()
	day05.NewStringWriter()
	day05.TestDivide()
	day06.GoRe()

	day04.RangeTest()

}

func test1() {
	guoguo := "guoguo"
	fmt.Println(guoguo)
	x := 100
	fmt.Println(x)
	y := 200
	fmt.Println(y)
	var guoguo1 string = "guoguo1"
	fmt.Println(guoguo1)
}

func testIota() {
	fmt.Println(a, b, c)
}

func testIosta2() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func testLeft() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

func testBool() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("a && b 为 true")
	}

	if a || b {
		fmt.Println("a || b 为 true")
	}

	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

func testFor() {
	var a int = 10
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

func testLoop() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}

}

func testBreakLoop() {
outerLoop:
	for i := 0; i < 2; i++ {
	innerLoop:
		for j := 0; j < 5; j++ {
			if j == 4 {
				break outerLoop
			}
			if j == 2 {
				break innerLoop
			}
			fmt.Printf("i的值为 : %d, j的值为 : %d\n", i, j)
		}
	}
}

func testBreak2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Received from ch1.")
	case <-ch2:
		fmt.Println("Received from ch2.")
		break // 跳出 select 语句
	}
}

func testSwitch() {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
