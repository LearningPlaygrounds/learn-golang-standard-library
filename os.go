package main

import (
	"fmt"
	"os"
)

// Package OS pada golang merupakan package yang digunakan untuk mengakses sistem yang sedang digunakan
func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
