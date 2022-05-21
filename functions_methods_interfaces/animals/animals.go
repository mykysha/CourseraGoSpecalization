package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	errUnknownCommand        = errors.New("command unknown")
	errUnknownAnimalType     = errors.New("animal type unknown")
	errWrongNumberOfCommands = errors.New("wrong number of commands")
)

type Animal struct {
	animalType string
	food       string
	locomotion string
	noise      string
	writer     *bufio.Writer
}

func (a *Animal) Init(aT, f, l, n string) {
	a.animalType = aT
	a.food = f
	a.locomotion = l
	a.noise = n
}

func (a Animal) Eat() error {
	err := writer(a.food+"\n", a.writer)
	if err != nil {
		return fmt.Errorf("eat: %w", err)
	}

	return nil
}

func (a Animal) Move() error {
	err := writer(a.locomotion+"\n", a.writer)
	if err != nil {
		return fmt.Errorf("move: %w", err)
	}

	return nil
}

func (a Animal) Speak() error {
	err := writer(a.noise+"\n", a.writer)
	if err != nil {
		return fmt.Errorf("speak: %w", err)
	}

	return nil
}

func (a Animal) Serve(c string) error {
	if c == "eat" {
		err := a.Eat()
		if err != nil {
			return fmt.Errorf("eat: %w", err)
		}

		return nil
	}

	if c == "move" {
		err := a.Move()
		if err != nil {
			return fmt.Errorf("serve: %w", err)
		}

		return nil
	}

	if c == "speak" {
		err := a.Speak()
		if err != nil {
			return fmt.Errorf("serve: %w", err)
		}

		return nil
	}

	return fmt.Errorf("serve: %w", errUnknownCommand)
}

func main() {
	var (
		w = bufio.NewWriter(os.Stdout)
		r = bufio.NewReader(os.Stdin)
	)

	cow := Animal{
		animalType: "cow",
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
		writer:     w,
	}

	bird := Animal{
		animalType: "bird",
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
		writer:     w,
	}

	snake := Animal{
		animalType: "snake",
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
		writer:     w,
	}

	zoo := make(map[string]Animal)

	zoo["cow"] = cow
	zoo["bird"] = bird
	zoo["snake"] = snake

	manager(zoo, w, r)
}

func manager(zoo map[string]Animal, w *bufio.Writer, r *bufio.Reader) {
	for {
		err := writer(">", w)
		if err != nil {
			log.Println("could not read write prompt, trying again")

			continue
		}

		line, err := reader(r)
		if err != nil {
			log.Println("could not read your command")

			continue
		}

		animalType, action, err := commandsConstructor(line)
		if err != nil {
			log.Println(errors.Unwrap(err))

			continue
		}

		err = zoo[animalType].Serve(action)
		if err != nil {
			log.Printf("could not write %s information (action %s)", animalType, action)

			continue
		}
	}
}

func commandsConstructor(s string) (string, string, error) {
	commands := strings.Split(s, " ")

	if wantedCommands := 2; len(commands) != wantedCommands {
		return "", "", fmt.Errorf("command: %w", errWrongNumberOfCommands)
	}

	n := commands[0]
	a := commands[1]

	err := checkCompliance(n, a)
	if err != nil {
		return "", "", fmt.Errorf("command: %w", err)
	}

	return n, a, nil
}

func checkCompliance(n, a string) error {
	if n != "cow" && n != "bird" && n != "snake" {
		return fmt.Errorf("check: %w", errUnknownAnimalType)
	}

	if a != "move" && a != "eat" && a != "speak" {
		return fmt.Errorf("check: %w", errUnknownCommand)
	}

	return nil
}

func writer(msg string, w *bufio.Writer) error {
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
