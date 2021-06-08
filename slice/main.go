package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func main() {
	// khai bao slice 
	var slice []int
	fmt.Println(slice)
	// khai bao va khoi tao gia tri
	var slice1 = []int {1,2,3,4}
	fmt.Println(slice1)
	// tao mot slice tu 1 array 
	var array = [4]int {1,2,3,4}
	slice2 := array[1:3] //array[1] -> array[3-1] <=> array[2]
	fmt.Println(slice2)

	slice3 := array[:] // get all value from array 
	fmt.Println(slice3)
	fmt.Println("========================")
	slice4 := array[2:]
	fmt.Println(slice4)
	fmt.Println("========================")
	slice5 := array[:3]
	fmt.Println(slice5)

	// reference type 
	var array1 = [5]int {1,2,3,4,5}
	slice9 := array1[:]
	slice9[0] = 777;
	fmt.Println(slice9)
	fmt.Println(array1)
	// length and capcity cua slice 
	counties := [...]string {"VN","USA","UK","CHINA"}
	slice10 := counties[2:3]
	fmt.Println(slice10)
	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))
	// len: so luong phan tu cua slice
	// cap: so luong phan tu cua underflying array bat tu vi tri start khi ma slice duoc tao
	fmt.Println("====================")
	//make 
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))
	// append 
	fmt.Println("====================")
	var slice12 []int
	slice13 := append(slice12, 100)
	fmt.Println(slice13)
	//copy 
	src := []string {"A","B","C","D"}
	dest := make([]string, 2)
	copy(dest, src)
	fmt.Println(dest)
	// delete item of array
	src1 := append(src[:1], src[2:]...)
	fmt.Println(src1)
}
