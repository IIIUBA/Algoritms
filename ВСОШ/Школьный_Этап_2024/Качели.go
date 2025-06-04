package main

import (
	"fmt"
	"math"
)

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func isPutStone(a, b, c, d int) int {

	reqA := 0
	diff := int(math.Abs(float64(a + b - c)))
	if d < diff {
		reqA = diff - d
	} else {
		reqA = 0
	}

	reqB := 0
	diff = int(math.Abs(float64(a + c - b)))
	if d < diff {
		reqB = diff - d
	} else {
		reqB = 0
	}

	reqC := 0
	diff = int(math.Abs(float64(b + c - a)))
	if d < diff {
		reqC = diff - d
	} else {
		reqC = 0
	}

	return min(reqA, reqB, reqC)
}

func main() {
	var a, b, c, d int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	_, err = fmt.Scan(&c)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	_, err = fmt.Scan(&d)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Подходящий камень:", isPutStone(a, b, c, d))
}
