package main

import "fmt"

func main() {
	var x float64
	fmt.Printf("Enter float:\t")

	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("could not read")

		return
	}
	fmt.Println("Int is: ", int(x))
}
