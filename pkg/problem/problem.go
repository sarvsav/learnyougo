package problem

import (
	"github.com/sarvsav/learnyougo/pkg/render"
)

func Problem() bool {
	render.Render("./exercises/1/problem.en.md")
	return true
}
