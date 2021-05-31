package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	var myBool bool = true // false
	fmt.Println(myBool)

	var myString string = "hello"
	fmt.Println(myString)
	// int 
	var myInt int = 10
	fmt.Println(myInt)

	// int8, int32, int64
	//1.Range int8 
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)
	//1.Range int16 
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)
	//1.Range int32
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)
	//1.Range int64
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)
	//2.bits 
	fmt.Println("=========================")
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))
	fmt.Println("=========================")
	//uint 
	var myUInt uint = 10
	fmt.Println(myUInt)
	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	// float 
	fmt.Println("=========================")
	var myFloat float32 = 10.01
	fmt.Println(myFloat)
	// complex 
	//Rune 
	var myString1 string = "Nhẫn"
	runes := []rune(myString1)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	//
	fmt.Println("================================")
	var myRune rune = 'A'
	fmt.Printf("%T", myRune)
	fmt.Println("================================")
	var a int = 1
	var b float64 = 2.1
	fmt.Println(float64(a) + b)
	const PI = 3.4
	fmt.Println(PI)
}