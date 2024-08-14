package day03

import "fmt"

func TestArray() {
	var n [10]int
	var j, i int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d]=%d\n", j, n[j])
	}
}

func Balance() {
	var i, j, k int
	balance := [5]float32{100.0, 2.0, 3.4, 7.0, 50.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d]=%d\n", j, balance[i])
	}

	balance2 := [...]float32{100.0, 2.0, 3.4, 7.0, 50.03, 433.2}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance2[%d]=%d\n", j, balance2[i])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.011111}

	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d]=%d\n", k, balance3[k])
	}
}

func VariaArray() {
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{3, 4, 5}

	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Println("first element:")
	fmt.Println(values[0][0])

}

func ManyString() {
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}

func GetAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
