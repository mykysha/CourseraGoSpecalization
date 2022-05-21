package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"sync"
)

func main() {
	numberOfRoutines := 4

	// sliceLen := 100
	// maxValue := 1000
	// slice := createRandomSlice(sliceLen, maxValue)

	slice, sliceLen := readSlice()

	_, _ = os.Stdout.WriteString(fmt.Sprintf("Unsorted slice: %v\n", slice))

	subSlices := splitSlice(slice, numberOfRoutines) // separating slice into numberOfRoutines parts
	wg := new(sync.WaitGroup)

	for i := 0; i < numberOfRoutines; i++ {
		wg.Add(1)

		ind := i

		go func(wg *sync.WaitGroup) { // starting numberOfRoutines goroutines to sort subslices
			_, _ = os.Stdout.WriteString(fmt.Sprintf("Sub slice that is sorted by routine №%d: %v\n",
				ind, subSlices[ind]))

			sortSlice(subSlices[ind])

			wg.Done()
		}(wg)
	}

	wg.Wait()

	sortedSlice := make([]int, sliceLen)

	for i := 0; i < sliceLen; i++ {
		minValue := math.MaxInt
		minSliceInd := -1

		// combining sorted slices
		for j := 0; j < numberOfRoutines; j++ {
			if len(subSlices[j]) < 1 {
				continue
			}

			if subSlices[j][0] <= minValue {
				minValue = subSlices[j][0]
				minSliceInd = j
			}
		}

		subSlices[minSliceInd] = subSlices[minSliceInd][1:]

		sortedSlice[i] = minValue
	}

	// printing results
	if !sort.IntsAreSorted(sortedSlice) {
		_, _ = os.Stdout.WriteString("Slice is not sorted!")
	} else {
		_, _ = os.Stdout.WriteString(fmt.Sprintf("Sorted slice: %v\n", sortedSlice))
	}
}

func sortSlice(slice []int) {
	sort.Ints(slice)
}

func readSlice() ([]int, int) {
	_, _ = os.Stdout.WriteString("Input int slice in line ending by any non-number symbol like" +
		" \"1 4 2 0 -9 67 4 7 6 2 3 4 5 6 7 88 9 STOP\"\n")

	var slice []int

	for {
		var n int

		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		slice = append(slice, n)
	}

	return slice, len(slice)
}

func createRandomSlice(sliceLen, maxValue int) []int {
	slice := make([]int, sliceLen)

	for i := 0; i < sliceLen; i++ {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(maxValue)))

		slice[i] = int(val.Int64())
	}

	return slice
}

func splitSlice(slice []int, numberOfParts int) [][]int {
	subSlices := make([][]int, numberOfParts)

	for i := 0; i < numberOfParts; i++ {
		subSlices[i] = slice[i*len(slice)/numberOfParts : (i+1)*len(slice)/numberOfParts]
	}

	return subSlices
}
