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
	r                        = bufio.NewReader(os.Stdin)
	w                        = bufio.NewWriter(os.Stdout)
	errUnknownCommand        = errors.New("command unknown")
	errEmptyCommand          = errors.New("found no command")
	errUnknownAnimalType     = errors.New("animal type unknown")
	errWrongNumberOfCommands = errors.New("wrong number of commands")
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	c.food = "grass"

	if err := writer(c.food + "\n"); err != nil {
		log.Println(err)
	}
}

func (c Cow) Move() {
	c.locomotion = "walk"

	if err := writer(c.locomotion + "\n"); err != nil {
		log.Println(err)
	}
}

func (c Cow) Speak() {
	c.noise = "moo"

	if err := writer(c.noise + "\n"); err != nil {
		log.Println(err)
	}
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	b.food = "worms"

	if err := writer(b.food + "\n"); err != nil {
		log.Println(err)
	}
}

func (b Bird) Move() {
	b.locomotion = "fly"

	if err := writer(b.locomotion + "\n"); err != nil {
		log.Println(err)
	}
}

func (b Bird) Speak() {
	b.noise = "peep"

	if err := writer(b.noise + "\n"); err != nil {
		log.Println(err)
	}
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	s.food = "mice"

	if err := writer(s.food + "\n"); err != nil {
		log.Println(err)
	}
}

func (s Snake) Move() {
	s.locomotion = "slither"

	if err := writer(s.locomotion + "\n"); err != nil {
		log.Println(err)
	}
}

func (s Snake) Speak() {
	s.noise = "hsss"

	if err := writer(s.noise + "\n"); err != nil {
		log.Println(err)
	}
}

func main() {
	zoo := make(map[string]Animal)

	for {
		line, err := prompt()
		if err != nil {
			log.Println(errors.Unwrap(err))
		}

		header, name, footer, err := commandsConstructor(line)
		if err != nil {
			log.Println(errors.Unwrap(err))
		}

		switch header {
		case "newanimal":
			if _, ok := zoo[name]; ok {
				err = writer("We already have an animal with this name. Choose another name.\n")
				if err != nil {
					log.Println(errors.Unwrap(err))
				}

				continue
			}

			newAnimal, err := createAnimal(footer)
			if err != nil {
				log.Println(errors.Unwrap(err))

				continue
			}

			err = writer("Created it!\n")
			if err != nil {
				log.Println(errors.Unwrap(err))
			}

			zoo[name] = newAnimal

		case "query":
			if _, ok := zoo[name]; !ok {
				err = writer("We do not have an animal with this name. Choose another name.\n")
				if err != nil {
					log.Println(errors.Unwrap(err))
				}

				continue
			}

			switch footer {
			case "eat":
				zoo[name].Eat()

			case "move":
				zoo[name].Move()

			case "speak":
				zoo[name].Speak()

			default:
				log.Println(errUnknownCommand)
			}

		default:
			log.Println(errUnknownCommand)
		}
	}
}

func createAnimal(animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{}, nil

	case "bird":
		return Bird{}, nil

	case "snake":
		return Snake{}, nil

	default:
		return nil, fmt.Errorf("creating animal: %w", errUnknownAnimalType)
	}
}

func prompt() (string, error) {
	if err := writer("> "); err != nil {
		return "", fmt.Errorf("prompting: %w", err)
	}

	c, err := reader()
	if err != nil {
		return "", fmt.Errorf("prompting: %w", err)
	}

	return c, nil
}

func commandsConstructor(s string) (string, string, string, error) {
	var animal, command string

	commands := strings.Split(s, " ")
	numberOfCommands := len(commands)

	if wanted := 3; numberOfCommands != wanted {
		return "", "", "", fmt.Errorf("command: %w", errWrongNumberOfCommands)
	}

	header := commands[0]
	name := commands[1]
	v2 := commands[2]

	switch header {
	case "newanimal":
		animal = v2

	case "query":
		command = v2

		if command == "" {
			return "", "", "", fmt.Errorf("command: %w", errEmptyCommand)
		}

	default:
		return "", "", "", fmt.Errorf("command: %w", errUnknownCommand)
	}

	err := checkCompliance(animal, command)
	if err != nil {
		return "", "", "", fmt.Errorf("command: %w", err)
	}

	return header, name, v2, nil
}

func checkCompliance(animalType, command string) error {
	if !typeChecker(animalType) {
		return fmt.Errorf("check: %w", errUnknownAnimalType)
	}

	if !commandChecker(command) {
		return fmt.Errorf("check: %w", errUnknownCommand)
	}

	return nil
}

func typeChecker(animal string) bool {
	possibleTypes := []string{"cow", "bird", "snake", ""}

	for _, val := range possibleTypes {
		if animal == val {
			return true
		}
	}

	return false
}

func commandChecker(command string) bool {
	possibleCommands := []string{"eat", "move", "speak", ""}

	for _, val := range possibleCommands {
		if command == val {
			return true
		}
	}

	return false
}

func writer(msg string) error {
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

func reader() (string, error) {
	v, _, err := r.ReadLine()
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}

	return string(v), nil
}
