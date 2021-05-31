package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	result := reverse(900000)
	fmt.Println(result)
}
func reverse(x int) int {
	var buffer bytes.Buffer
	if x == 0 {
		return 0
	}
	if x < 0 {
		var y float64 = math.Abs(float64(x)) 
		x = int(y)
		buffer.WriteString("-")
	}
	cardLeter := strconv.Itoa(x) 
	cardArrayLeter := strings.Split(cardLeter,"")
	length := len(cardArrayLeter) - 1
	
	var max int64 = 0
	for i := length ; i >= 0; i-- {
		intVal, _ := strconv.ParseInt(cardArrayLeter[i], 0, 32)
		if max < intVal {
			max = intVal
		}
		buffer.WriteString(strconv.FormatInt(intVal, 10))
	}
	s := buffer.String()
	reverseX, _ := strconv.Atoi(s)	
	if reverseX > int(math.Pow(2, 31)) || reverseX < int(math.Pow(-2, 31)) {
		return 0;
	}
	return reverseX
}
