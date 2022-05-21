package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	makejson()
}

func makejson() {
	info := make(map[string]string, 2)

	fmt.Print("\nenter name:\t")

	name, err := readLine()
	if err != nil {
		fmt.Println(err)
	}

	info["name"] = name

	fmt.Print("\nenter address:\t")

	addr, err := readLine()
	if err != nil {
		fmt.Println(err)
	}

	info["address"] = addr

	myJSON, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		fmt.Println("could not marshal: ", err)
	}

	fmt.Println("your JSON:\n\n", string(myJSON))
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	s, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read: %w", err)
	}

	s = strings.Trim(s, "\n")

	return s, nil
}
