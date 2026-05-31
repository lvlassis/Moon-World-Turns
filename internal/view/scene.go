package view

import (
	"github.com/g3n/engine/renderer"
)

type Scene interface {

	Render(renderer *renderer.Renderer)

	Update(deltaTime float64)

	OnResize(width, height int)
}
