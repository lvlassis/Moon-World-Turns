package game

import (
	"log"
	"os"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"gopkg.in/yaml.v3"
)


type Stage struct {
	stageData
	*core.Node
}

type stageData struct {
	StageName  string     `yaml:"name"`
	Ground     groundData `yaml:"ground"`
	Lights     []lightData `yaml:"lights"`
}

type groundData struct {
	Image string `yaml:"image"`
	Color string `yaml:"color"`
}

type lightData struct {
	Type  string  `yaml:"type"`
	Color struct {
		Red   float32 `yaml:"red"`
		Green float32 `yaml:"green"`
		Blue  float32 `yaml:"blue"`
	}  
	Intensity float32 `yaml:"intensity"`
	Position struct {
		X float32 `yaml:"x"`
		Y float32 `yaml:"y"`
		Z float32 `yaml:"z"`
	}
}



func LoadStage(path string) *Stage {
	// Lê o arquivo
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao ler Stage %s: %s", path, err)
	}

	var stageData stageData
	err = yaml.Unmarshal(data, &stageData)
	if err != nil {
		log.Fatalf("Erro ao parsear YAML de %s: %s", path, err)
	}

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


	return &Stage{
		stageData: stageData,
		Node:      stageNode,
	}
}
