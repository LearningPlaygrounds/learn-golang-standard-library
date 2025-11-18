package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Mengubah string menjadi boolean
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Boolean:", result)
	}

	// Mengubah string menjadi integer
	intResult, err := strconv.ParseInt("12345", 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Integer:", intResult)
	}

	// Mengubah string menjadi float
	floatResult, err := strconv.ParseFloat("3.14159", 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Float:", floatResult)
	}

	// Mengubah integer menjadi string
	strResult := strconv.Itoa(6789)
	fmt.Println("Integer to String:", strResult)

	// Mengubah float menjadi string
	floatStrResult := strconv.FormatFloat(2.71828, 'f', 2, 64)
	fmt.Println("Float to String:", floatStrResult)
}
