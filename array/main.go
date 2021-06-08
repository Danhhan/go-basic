package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func main() {
	// khai bao array
	var myArray [4]int 	
	fmt.Println(myArray)
	// gan gia tri
	myArray[0] = 123
	fmt.Println(myArray)
	// khai bao co khoi tao gia tri
	array := [3]int {1,2,3}
	fmt.Println(array)
	// size array 
	fmt.Println(len(array))
	//khai bao mang khong can set size 
	array3 := [...]string {"Vinfast","Huynhdai"}
	fmt.Println(array3)
	// cach 1
	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}
	// cach 2
	for index, value := range array3 {
		fmt.Printf("i=%d value=%s", index, value)
		fmt.Println()
	}
}
