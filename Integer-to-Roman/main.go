package main

import (
	"bytes"
	"fmt"
)

func main() {
	result := convertToRoman(4)
	fmt.Println(result)
}
func convertToRoman(x int) string {
	var buffer bytes.Buffer
	if x >= 1 || x < 3 {	
		for i := 1; i <= x; i++ {
			buffer.WriteString("I")	
		}
	}
	return buffer.String()
}

