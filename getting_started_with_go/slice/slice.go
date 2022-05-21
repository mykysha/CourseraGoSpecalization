package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	mySlice := make([]int, 3)

	for {
		newValue, exit := reader()
		if exit {
			break
		}

		mySlice = sorter(mySlice, newValue)
	}

	fmt.Println("Program ended.")
}

func reader() (int, bool) {
	for {
		var (
			input string
			val   int
		)

		fmt.Print("Enter number:\t")

		_, err := fmt.Scan(&input)
		if err != nil {
			continue
		}

		if input == "X" {
			return 0, true
		}

		val, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("inputted text is not int64")
			continue
		}

		return val, false
	}
}

func sorter(mySLice []int, val int) []int {
	mySLice = append(mySLice, val)

	sort.Ints(mySLice)

	fmt.Println("sorted:\t", mySLice)

	return mySLice
}
