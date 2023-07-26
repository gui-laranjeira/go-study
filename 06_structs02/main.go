// Create a program that reads the contents of a text file then prints its contents to the terminal
// The file to open should be provided as an arugmento to the program when it is executed at the terminal. For example, 'go run main.go myfile.txt'.
// To read in the arguments provided to a program, you can reference the variable os.Args, whicch is a slice of tyhpe string

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
