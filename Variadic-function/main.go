package main

import "fmt"

// variadic funcion 
// 1. Khai bao variadic func
// 2. Pass 1 slice vao variadic func
// 3. Change slice 
func addItem(item int, list ...int) {
	//100,200,300,400 => int[] {100,200,300,400}
	list = append(list, item)
	fmt.Println(list)
}
func main() {
	addItem(1, 100,200,300,400)
	var list = []int {1,2,3,4}   
	addItem(100, list...)
	change(list...);
	fmt.Println(list)
}
func change(list ...int) {
	list[0] = 999;
}