package mypack

import (
	"fmt"
	"math"
)

func Add(i int, j int) int {
	fmt.Println("Just a fix and commiting it")
	return i + j + 3
}

func Pow(i float64, j float64) int {
	fmt.Println("Just a fix and commiting it")
	return int(math.Pow(i, j))
}
