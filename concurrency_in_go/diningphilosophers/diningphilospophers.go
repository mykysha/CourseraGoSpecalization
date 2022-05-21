package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type host struct {
	startChan     chan interface{}
	allowanceChan chan interface{}
}

func newHost(numberLimit int) *host {
	startChan := make(chan interface{})
	allowanceChan := make(chan interface{}, numberLimit)

	return &host{
		startChan:     startChan,
		allowanceChan: allowanceChan,
	}
}

func (h host) hostRoutine() {
	for {
		<-h.startChan
		h.allowanceChan <- nil
	}
}

func (h host) startEating() {
	h.startChan <- nil
}

func (h host) hostStoppedEating() {
	<-h.allowanceChan
}

type philosopher struct {
	name         string
	eatTimes     int
	thinkTime    time.Duration
	eatingTime   time.Duration
	dominantHand *sync.Mutex
	otherHand    *sync.Mutex
}

func main() {
	numberOfPhilosophers := 5

	ph := make([]philosopher, numberOfPhilosophers)

	for i := 0; i < numberOfPhilosophers; i++ {
		ph[i] = philosopher{
			name:         strconv.Itoa(i + 1),
			eatTimes:     3,
			thinkTime:    time.Second,
			eatingTime:   time.Second,
			dominantHand: new(sync.Mutex),
			otherHand:    new(sync.Mutex),
		}
	}

	host := newHost(2)

	go host.hostRoutine()

	dining := new(sync.WaitGroup)

	dining.Add(5)

	fork0 := new(sync.Mutex)
	forkLeft := fork0

	for i := 1; i < len(ph); i++ {
		forkRight := new(sync.Mutex)

		ph[i].dominantHand = forkRight
		ph[i].otherHand = forkLeft

		go diningProblem(&ph[i], dining, host)

		forkLeft = forkRight
	}

	go diningProblem(&ph[0], dining, host)

	dining.Wait()
}

func diningProblem(ph *philosopher, dining *sync.WaitGroup, host *host) {
	for h := ph.eatTimes; h > 0; h-- {
		ph.dominantHand.Lock()
		ph.otherHand.Lock()

		host.startEating()

		write(fmt.Sprintf("starting to eat %s\n", ph.name))

		time.Sleep(ph.eatingTime)

		host.hostStoppedEating()

		write(fmt.Sprintf("finishing eating %s\n", ph.name))

		ph.dominantHand.Unlock()
		ph.otherHand.Unlock()

		time.Sleep(ph.thinkTime)
	}

	dining.Done()
}

func write(data ...any) {
	_, _ = os.Stdout.WriteString(fmt.Sprint(data...))
}
