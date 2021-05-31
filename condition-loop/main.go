package main

import "fmt"

func main() {
	number := 10 

	if number == 10 {
		fmt.Println("number = ", number)
	}

	if number < 10 {
		fmt.Println("Number < 10")
	} else {
		fmt.Println("Number = 10")
	}

	if a := 100; a > 100 {
		fmt.Println("a > 100")
	} else {
		fmt.Println("a = 100")
	}

	// switch case
	number1 := 10
	switch number1 {
	case 1:
		fmt.Println("Number = 1")
	case 2:
		fmt.Println("Number = 2")
	case 3:
		fmt.Println("Number = 3")
	case 4:
		fmt.Println("Number = 4")
	case 5:
		fmt.Println("Number = 5")
	case 6:
		fmt.Println("Number = 6")
	default: 
		fmt.Println("Unknow")
	}
	//FallThrou
	fmt.Println("=========FallThrough")
	number2 := 10
	switch number2 {
	case 10:
		if number2 == 10 {
			fmt.Println("Number = 10")
			break
		}
		fallthrough
	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 3:
		fmt.Println("Number = 3")
		fallthrough
	case 4:
		fmt.Println("Number = 4")
		fallthrough
	case 5:
		fmt.Println("Number = 5")
		fallthrough
	case 6:
		fmt.Println("Number = 6")
		fallthrough
	default: 
		fmt.Println("Unknow")
	}
	// loop
	fmt.Println("Loop =======================")
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(i)
	} 
	fmt.Println("out of loop")
	// while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	//infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }
	// loop multi condition
	for i, j := 1, 2; i< 10 && j < 10; i, j = i + 1, j + 1 {
		fmt.Println(i)
		fmt.Println(i)
	}
}