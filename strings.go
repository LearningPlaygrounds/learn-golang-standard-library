package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Thomas Alberto", "Thomas"))
	fmt.Println(strings.Split("Thomas Alberto", " "))
	fmt.Println(strings.ToUpper("Thomas Alberto"))
	fmt.Println(strings.ToLower("Thomas Alberto"))
	fmt.Println(strings.HasPrefix("Thomas Alberto", "Tho"))
	fmt.Println(strings.HasSuffix("Thomas Alberto", "Alberto"))
	fmt.Println(strings.Index("Thomas Alberto", "Alberto"))
	fmt.Println(strings.ReplaceAll("Thomas Alberto", "Alberto", "Smith"))
}
