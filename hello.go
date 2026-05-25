package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"log"
	"time"
)

func NewCharacter(id string, position int, image_path string) *graphic.Sprite {
	char_texture, err := texture.NewTexture2DFromImage(image_path)
	if err != nil {
		log.Fatal("Failed to load assets/ness.png: ", err)
	}
	var proportion float32 = float32(char_texture.Height()) / float32(char_texture.Width())
	
	char_texture.SetMagFilter(gls.NEAREST)
	char_texture.SetMinFilter(gls.NEAREST)
	// Create material with the texture
	char_material := material.NewStandard(math32.NewColor("white"))
	char_material.AddTexture(char_texture)
	char_material.SetTransparent(true)
	// Create sprite using the proper Sprite type
	char_sprite := graphic.NewSprite(1, proportion, char_material)

	if position == 1 {
		char_sprite.SetPosition(-1, proportion/2, 0)
	} else if position == 2 {
		char_sprite.SetPosition(1, proportion/2, 0)
	}
	return char_sprite
}

func DesenhaCena(scene *core.Node) {
	// Cria chão
	chaoTex, err := texture.NewTexture2DFromImage("assets/grass.png")
	if err != nil {
		log.Fatal("Failed to load assets/grass.png: ", err)
	}
	chaoTex.SetRepeat(8, 8)
	chaoTex.SetWrapS(gls.REPEAT)
	chaoTex.SetWrapT(gls.REPEAT)
	chaoTex.SetMagFilter(gls.NEAREST)
	chaoTex.SetMinFilter(gls.NEAREST)
	chao := geometry.NewPlane(10, 10)
	chao.ApplyMatrix(math32.NewMatrix4().MakeRotationX(-math32.Pi / 2))
	chaoMat := material.NewStandard(math32.NewColor("DarkGreen"))
	chaoMat.AddTexture(chaoTex)
	chaoMesh := graphic.NewMesh(chao, chaoMat)
	scene.Add(chaoMesh)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

}

func main() {
	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 1, 6)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Set up callback to close the app when Escape is pressed
	a.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		kev := ev.(*window.KeyEvent)
		if kev.Key == window.KeyEscape || kev.Key == window.KeyQ {
			a.Exit()
		}
	})

	// Controlar a câmera
	a.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		var speed float32 = 0.1
		kev := ev.(*window.KeyEvent)
		switch kev.Key {
		case window.KeySpace:
			cam.SetPosition(cam.Position().X, cam.Position().Y+speed, cam.Position().Z)
		case window.KeyLeftShift:
			cam.SetPosition(cam.Position().X, cam.Position().Y-speed, cam.Position().Z)
		case window.KeyW:
			cam.SetPosition(cam.Position().X, cam.Position().Y, cam.Position().Z-speed)
		case window.KeyS:
			cam.SetPosition(cam.Position().X, cam.Position().Y, cam.Position().Z+speed)
		case window.KeyA:
			cam.SetPosition(cam.Position().X-speed, cam.Position().Y, cam.Position().Z)
		case window.KeyD:
			cam.SetPosition(cam.Position().X+speed, cam.Position().Y, cam.Position().Z)
		}
	})

	// Cria Char 1
	// Load Ness sprite texture
	char1 := NewCharacter("charmander", 1, "assets/charmander.png")
	scene.Add(char1)

	char2 := NewCharacter("squirtle", 2, "assets/squirtle.png")
	scene.Add(char2)

	DesenhaCena(scene)

	// Set background color to gray
	a.Gls().ClearColor(0.8, 0.8, 0.8, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}
