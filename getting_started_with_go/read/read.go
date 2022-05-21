package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type names struct {
	Fname string
	Lname string
}

func main() {
	var (
		filename string
		answ     []names
	)
	fmt.Println("Write filename:")

	_, err := fmt.Scan(&filename)
	if err != nil {
		log.Println(err)
	}

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}

	lines := strings.Split(string(fileData), "\n")

	for _, val := range lines {
		fullName := strings.Split(val, " ")

		s := names{Fname: fullName[0], Lname: fullName[1]}
		answ = append(answ, s)
	}

	for _, val := range answ {
		fmt.Printf("first name: %s\t last name: %s\n", val.Fname, val.Lname)
	}
}
