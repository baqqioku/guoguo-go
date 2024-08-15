package day04

import (
	"fmt"
)

func Swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 函数接受一个数组作为参数
func ModifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// 函数接受一个数组指针作为参数
func ModifyArrayWithPointer(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}

}

func Pptr() {
	var a int
	var ptr *int
	var pptr **int
	a = 300
	ptr = &a
	pptr = &ptr
	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Println(Books{title: "go 语言", author: "www.runoob.com"})
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func PrintBook(book *Books) {
	fmt.Println(book)
}

func SlicelStr() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 8}

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func MakeSlice() {
	nums := make([]int, 3, 5)

	fmt.Println("Initial slice:", nums) // Outputs: [0 0 0]

	// Modifying elements within the slice's length
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30

	fmt.Println("Modified slice:", nums) // Outputs: [10 20 30]

	// Appending elements within the capacity
	nums = append(nums, 40)
	nums = append(nums, 50)

	fmt.Println("Slice after append:", nums)         // Outputs: [10 20 30 40 50]
	fmt.Println("Length after append:", len(nums))   // Outputs: 5
	fmt.Println("Capacity after append:", cap(nums)) // Outputs: 5

	// Appending beyond capacity
	nums = append(nums, 60) // This will trigger a resize of the underlying array

	fmt.Println("Slice after exceeding capacity:", nums)         // Outputs: [10 20 30 40 50 60]
	fmt.Println("Length after exceeding capacity:", len(nums))   // Outputs: 6
	fmt.Println("Capacity after exceeding capacity:", cap(nums)) // Outputs: 10 or more
}
