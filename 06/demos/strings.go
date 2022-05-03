package main

import "fmt"

func main() {
	ourString := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%q\n", ourString[i])
	}

	otherString := "Hello World!"
	fmt.Println("5 index value:", otherString[5])
	fmt.Println("substring:", otherString[0:5])
}
