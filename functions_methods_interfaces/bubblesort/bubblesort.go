package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var errMuchNumbers = errors.New("you have inserted more numbers than needed")

func main() {
	var (
		w = bufio.NewWriter(os.Stdout)
		r = bufio.NewReader(os.Stdin)
	)

	msg := "ENTER INTS DIVIDED BY SPACES (NOT MORE THAN 10)\n"

	err := write(w, msg)
	if err != nil {
		log.Println("First write: ", err)
	}

	intSlice, err := readSlice(r)
	if err != nil {
		log.Println("Slice read: ", err)
	}

	BubbleSort(intSlice)

	for _, val := range intSlice {
		p := strconv.Itoa(val) + " "

		err = write(w, p)
		if err != nil {
			log.Println("Results write: ", err)
		}
	}
}

func BubbleSort(n []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				swapped = true

				Swap(n, i)
			}
		}
	}
}

func Swap(n []int, pos int) {
	n[pos], n[pos+1] = n[pos+1], n[pos]
}

func write(w *bufio.Writer, msg string) error {
	out := []byte(msg)

	_, err := w.Write(out)
	if err != nil {
		return fmt.Errorf("write err: %w", err)
	}

	err = w.Flush()
	if err != nil {
		return fmt.Errorf("flush err: %w", err)
	}

	return nil
}

func readSlice(r *bufio.Reader) ([]int, error) {
	goalSize := 10

	returnSlice := make([]int, 0, goalSize)

	in, _, err := r.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("command line read: %w", err)
	}

	s := string(in)

	nums := strings.Split(s, " ")

	if len(nums) > goalSize {
		return nil, fmt.Errorf("slice separation: %w", errMuchNumbers)
	}

	for _, value := range nums {
		n, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("slice conversion: %w", err)
		}

		returnSlice = append(returnSlice, n)
	}

	return returnSlice, nil
}
