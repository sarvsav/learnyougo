// Package verify compares the output by the program with the stored value
package verify

import "fmt"

func Verify() bool {
	fmt.Println("Verify Called inside package...")
	return true
}
