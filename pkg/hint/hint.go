package hint

import "github.com/sarvsav/learnyougo/pkg/render"

func Hint() bool {
	render.Render("./exercises/1/hint.en.md")
	return true
}
