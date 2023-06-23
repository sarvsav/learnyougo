package hint

import "github.com/sarvsav/learnyougo/pkg/render"

// Hint() is a function that returns a boolean value
func Hint() bool {
	render.Render("./exercises/1/hint.en.md")
	return true
}
