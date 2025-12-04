package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Predefined animals
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var animalName, action string

		fmt.Print("> ") // prompt
		_, err := fmt.Scan(&animalName, &action)

		if err != nil {
			fmt.Println("Input error, try again.")
			continue
		}

		animal, ok := animals[animalName]
		if !ok {
			fmt.Println("Unknown animal")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Unknown action")
		}
	}
}
