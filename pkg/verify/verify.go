// Package verify compares the output by the program with the stored value
package verify

import (
	"fmt"
	"os"

	exercises "github.com/sarvsav/learnyougo/exercises/1"
)

func Verify(fileName string) bool {
	fmt.Println("Running and verifying the file:", fileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("File doesn't exist: %s\n", fileName)
		return false
	}
	if exercises.Output(fileName) {
		fmt.Println("Congratulations! Your output is correct.")
		return true
	}
	return false
}
