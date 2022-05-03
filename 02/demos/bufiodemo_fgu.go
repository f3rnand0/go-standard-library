package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//one()
	//two()
	three()
}

func one() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("Quitting")
			os.Exit(3)
		} else {
			fmt.Println("You typed " + scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func two() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("You typed " + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func three() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: filemaker <input file>")
	} else {
		fmt.Println("How would you like to see the text?")
		fmt.Println("1: ALL CAPS")
		fmt.Println("2: Title case")
		fmt.Println("3: Lower case")

		var option int
		_, err := fmt.Scanf("%d", &option)

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			switch option {
			case 1:
				fmt.Println(strings.ToUpper(scanner.Text()))
			case 2:
				fmt.Println(strings.Title(scanner.Text()))
			case 3:
				fmt.Println(strings.ToLower(scanner.Text()))
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
