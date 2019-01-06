package main

import "fmt"

func main() {
	months := [...]string{1: "January", 12: "December"}
	fmt.Printf("len : %d, %v\n", len(months), months)

	months2 := [...]string{"Jan", "Feb", "Nov", "Dec"}
	fmt.Printf("len : %d, %v\n", len(months2), months2)

	months3 := months2[0:2]
	fmt.Printf("len : %d, %v\n", len(months3), months3)
}
