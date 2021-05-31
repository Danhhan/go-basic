package main

import (
	"fmt"
)

func main() {
	var number int
	number = 10
	fmt.Println(number)
	var number1 int = 11
	fmt.Println(number1)

	// type inference 
	var email = "danhh@antoree.com"
	fmt.Println(email)

	// khai bao nhieu bien 
	//1. khai bao nhieu bien cung kieu du lieu
	var a, b int 
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)
	var a1,b1 int = 10, 11
	fmt.Println(a1)
	fmt.Println(b1)
	//2. khai bao nhieu bien khac kieu du lieu
	var (
		name string
		address string 
		age int 
	)
	name = "Danh"
	address = "viet nam"
	age = 25
	fmt.Println(name,address,age)
	//
	var (
		name1 string = "Danh1"
		address1 string = "Address1"
		age1 int = 25
	)
	fmt.Println(name1,address1,age1)
	// 
	var name2, address2, age2 = "Name2", "Address2", "Age2"
	fmt.Println(name2,address2,age2)

	language := "Golang"
	fmt.Println(language)
}