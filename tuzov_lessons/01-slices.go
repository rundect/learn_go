package main

import (
	"fmt"
)

func main(){
	base_array := [6]string{"f", "f", "@", "#", "z", "x"}
	var slice []string = base_array[:3]
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(base_array), cap(base_array), base_array)

	slice = append(slice, "Q")
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(base_array), cap(base_array), base_array)

	slice = append(slice, "w")
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(base_array), cap(base_array), base_array)

	slice = append(slice, "e")
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(base_array), cap(base_array), base_array)

	slice = append(slice, "t")
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(base_array), cap(base_array), base_array)

	var list[]int
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	list = []int{}
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	list = append(list, 1)
	fmt.Println(list != nil)
	fmt.Println(len(list) == 1)

	list1 := make([]int, 5, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(list1), cap(list1), list1)

	list2 := []int{1, 2, 3, 4, 8}
	fmt.Println(double(list2))

	list3 := []int{1, 2, 3, 4}
	fmt.Println("before", list3)
	//handle(list3)
	handle(list3[:1])
	fmt.Println("after", list3)

	list4 := make([]int, 4, 5)
	list5 := append(list4, 1)
	list4[0] = 5
	list5[0] = 9

	fmt.Println(list4)
	fmt.Println(list5)
}

func double(nums []int) []int{
	res := make([]int, 0, len(nums))

	for _, num := range nums{
		res = append(res, num * 2)
	}
	return res
}

func handle(list []int){
	//list[1] = 123
	//_ = append(list, 5)
	list = append(list, 5)
	fmt.Println("append", list)
}