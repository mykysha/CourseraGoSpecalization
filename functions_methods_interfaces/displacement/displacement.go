package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		w = bufio.NewWriter(os.Stdout)
		r = bufio.NewReader(os.Stdin)
	)

	a, v, s0, err := communicator(w, r)
	if err != nil {
		log.Fatalf("command line read fTl error: %v", err)
	}

	f := GenDisplaceFn(a, v, s0)

	err = counter(f, w, r)
	if err != nil {
		log.Fatalf("calculating error: %v", err)
	}
}

func GenDisplaceFn(a, v, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v*t + s0
	}

	return fn
}

func counter(f func(float64) float64, w *bufio.Writer, r *bufio.Reader) error {
	var stop bool

	err := writer(w, "\nIf you want to stop, type in 'E'")
	if err != nil {
		return fmt.Errorf("stop prompt: %w", err)
	}

	for !stop {
		err = writer(w, "\n\nEnter time:\t")
		if err != nil {
			return fmt.Errorf("time prompt: %w", err)
		}

		t, isE, err := getValue(r)
		if err != nil {
			err = writer(w, "\n\nWARNING! WRONG INPUT\n\n")
			if err != nil {
				return fmt.Errorf("warning message: %w", err)
			}

			continue
		}

		if isE {
			err = writer(w, "\n\nSHUTTING DOWN\n\n")
			if err != nil {
				return fmt.Errorf("final message: %w", err)
			}

			break
		}

		answer := fmt.Sprintf("\nCurrent displacement:\t%v", f(t))

		err = writer(w, answer)
		if err != nil {
			return fmt.Errorf("answering: %w", err)
		}
	}

	return nil
}

func communicator(w *bufio.Writer, r *bufio.Reader) (float64, float64, float64, error) {
	err := writer(w, "\nHello and welcome to the displacement calculator")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("greeting: %w", err)
	}

	err = writer(w, "\n\tEnter acceleration value:\t")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("acceleration prompt: %w", err)
	}

	a, _, err := getValue(r)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("acceleration: %w", err)
	}

	err = writer(w, "\n\tEnter velocity value:\t")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("velocity prompt: %w", err)
	}

	v, _, err := getValue(r)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("velocity: %w", err)
	}

	err = writer(w, "\n\tEnter initial displacement value:\t")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("initial displacement prompt: %w", err)
	}

	s0, _, err := getValue(r)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("initial displacement: %w", err)
	}

	return a, v, s0, nil
}

func getValue(r *bufio.Reader) (float64, bool, error) {
	s, err := reader(r)
	if err != nil {
		return 0, false, fmt.Errorf("getting value: %w", err)
	}

	if s == "E" {
		return 0, true, nil
	}

	s = strings.ReplaceAll(s, ",", ".")

	floatSize := 64

	n, err := strconv.ParseFloat(s, floatSize)
	if err != nil {
		return 0, false, fmt.Errorf("conversion: %w", err)
	}

	return n, false, nil
}

func writer(w *bufio.Writer, msg string) error {
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

func reader(r *bufio.Reader) (string, error) {
	v, _, err := r.ReadLine()
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}

	return string(v), nil
}
