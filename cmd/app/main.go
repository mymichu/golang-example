package main

import (
	"fmt"
)

//Average Calculates the average of two numbers
func Average(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	average := Average(10,5);
	s := fmt.Sprintf("%f", average)
	fmt.Println("Hello, world. "+s)
}