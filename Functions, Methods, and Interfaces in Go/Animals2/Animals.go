package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal 接口：Eat / Move / Speak 都不带参数、不返回值
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// name -> Animal
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			fmt.Println("Please enter: newanimal/query <name> <type/info>")
			continue
		}

		cmd := parts[0]
		name := parts[1]
		arg := parts[2]

		switch cmd {
		case "newanimal":
			switch arg {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type")
				continue
			}
			fmt.Println("Created it!")

		case "query":
			a, ok := animals[name]
			if !ok {
				fmt.Println("Animal not found")
				continue
			}
			switch arg {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
				// speak
			case "speak":
				a.Speak()
			default:
				fmt.Println("Unknown info type")
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
