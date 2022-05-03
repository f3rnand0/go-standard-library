package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) == 0 {
		fmt.Println("Yo must enter arguments")
	} else {
		mealTotal, _ := strconv.ParseFloat(args[0], 32)
		tipAmount, _ := strconv.ParseFloat(args[1], 32)
		fmt.Printf("Your meal total is %.2f", mealTotal*tipAmount)
	}
}
