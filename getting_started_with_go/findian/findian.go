package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter string:\t")

	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("could not read: ", err)
	}

	s = strings.Trim(s, "\n")

	s = strings.ToLower(s)

	if strings.Contains(s, "a") && string(s[0]) == "i" && string(s[len(s)-1]) == "n" {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
