package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func main() {
	// hello()
	// result := greeting("Danh")
	// fmt.Println(result)

	// w, h, area := recInfo(10, 20)
	// fmt.Println("Width ", w)
	// fmt.Println("Height ", h)
	// fmt.Println("area ", area)
	w, h, isSquare := recInfo1(21, 20)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("Width ", w)
		fmt.Println("Height ", h)
	}
}
func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}
// Multiple return values
func recInfo(w, h int) (int, int, int ) {
	area := w * h
	return w, h, area
}
func recInfo1(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}