package metadata

import (
	"testing"
)

func TestProblemList(t *testing.T) {
	got, _ := ProblemList("../exercises/")
	want := "Hello World!"
	if got[0] != want {
		t.Errorf("got: %s, want: %s", got[0], want)
	}
}
