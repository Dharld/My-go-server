package main

import (
	"fmt"
	"math"
)


func add(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	for i := range 10 {
		fmt.Println(math.Pow(float64(i), 2))
	}
}