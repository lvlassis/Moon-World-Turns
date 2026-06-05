package scenes

import (
	"log"
	"lvlassis/moon-world-turns/internal/game"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type StageView struct {
	stageData *game.StageData
	*core.Node
}


func NewStageView(stageData *game.StageData) *StageView {
	stageNode := core.NewNode()

	// Instancia o chão
	chaoTex, err := texture.NewTexture2DFromImage(stageData.Ground.Image)
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
	chaoMat := material.NewStandard(math32.NewColor(stageData.Ground.Color))
	chaoMat.AddTexture(chaoTex)
	chaoMesh := graphic.NewMesh(chao, chaoMat)

	stageNode.Add(chaoMesh)

	// Instancia as luzes
	for _, lightData := range stageData.Lights {
		color := &math32.Color{lightData.Color.Red, lightData.Color.Green, lightData.Color.Blue} 
		intensity := lightData.Intensity
		switch lightData.Type {
		case "ambient":
			stageNode.Add(light.NewAmbient(color, intensity))
		case "point":
			pointLight := light.NewPoint(color, intensity)
			pointLight.SetPosition(lightData.Position.X, lightData.Position.Y, lightData.Position.Z)
			stageNode.Add(pointLight)
		default:
			log.Printf("Tipo de luz desconhecido: %s", lightData.Type)
		}
	}


	return &StageView{
		stageData: stageData,
		Node:      stageNode,
	}
}
