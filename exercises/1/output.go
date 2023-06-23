package exercises

import (
	"fmt"
	"os/exec"
	"strings"
)

func Output(fileName string) bool {
	output, err := exec.Command("go", "run", fileName).Output()
	if err != nil {
		panic(err) // write the output with ResponseWriter
	}

	return verifyOutput(string(output))
}

func verifyOutput(got string) bool {
	want := "Hello World!"
	if strings.TrimSuffix(got, "\n") == want {
		return true
	}
	fmt.Printf("Expected output: %s and got %s\n", want, got)
	return false
}
