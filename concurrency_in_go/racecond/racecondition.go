package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// This program demonstrates problems with race condition via running two goroutines N times.

const numberOfRacingFunctions = 2

const (
	thousand = 1000
	million  = 1000000
	billion  = 1000000000
)

func main() {
	// test cycle settings.
	startValue := 5
	numberOfRuns := thousand
	printStep := numberOfRuns / 10
	goroutineNumber := 8

	startTime := time.Now()

	resultMap := test(startValue, numberOfRuns, printStep, goroutineNumber)

	tookTime := time.Since(startTime)

	analysis := fmt.Sprintf("Out of %d runs:", numberOfRuns)

	for result, quantity := range resultMap {
		analysis += fmt.Sprintf("\n\t%d resulted in \"%d\"", quantity, result)
	}

	analysis += fmt.Sprintf("\n\ntests were running for %v", tookTime)

	_, err := os.Stdout.WriteString(analysis)
	if err != nil {
		log.Println(err)
	}
}

func test(startValue, numberOfRuns, printStep, goroutineNumber int) map[int]int {
	if goroutineNumber < 1 {
		goroutineNumber = 1
	}

	if goroutineNumber > million/2 {
		goroutineNumber = million / 2
	}

	if goroutineNumber > numberOfRuns {
		goroutineNumber = numberOfRuns
	}

	resultMap := make(map[int]int)
	mu := new(sync.Mutex)

	testSequence := func(runNumber int) {
		testWG := new(sync.WaitGroup)

		testWG.Add(numberOfRacingFunctions)

		testNumber := startValue

		go increment(&testNumber, testWG)
		go double(&testNumber, testWG)

		testWG.Wait()

		mu.Lock()

		resultMap[testNumber]++

		if runNumber%printStep == 0 {
			_, err := os.Stdout.WriteString(strconv.Itoa(testNumber) + "\n")
			if err != nil {
				log.Println(err)
			}
		}

		mu.Unlock()
	}

	runs := make([]int, 0, goroutineNumber)

	for i := 0; i < goroutineNumber; i++ {
		runs = append(runs, numberOfRuns/goroutineNumber)
	}

	for i := 0; i < numberOfRuns%goroutineNumber; i++ {
		runs[i]++
	}

	runWG := new(sync.WaitGroup)

	for i, workedOn := 0, 0; workedOn < numberOfRuns; i++ {
		runWG.Add(1)

		go func(start, finish int) {
			for j := start; j < finish; j++ {
				testSequence(j)
			}

			runWG.Done()
		}(workedOn, workedOn+runs[i])

		workedOn += runs[i]
	}

	runWG.Wait()

	return resultMap
}

func increment(a *int, wg *sync.WaitGroup) {
	*a++

	wg.Done()
}

func double(a *int, wg *sync.WaitGroup) {
	*a *= 2

	wg.Done()
}
