package main

import (
	"fmt"
	_ "image/png"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	img  *ebiten.Image
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.PressedKeys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	str := ""
	for _, k := range g.keys {
		str += fmt.Sprintf("%s ", k)
	}
	ebitenutil.DebugPrint(screen, "Hello, World!" + str)
	// screen.DrawImage(g.img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	img, _, err := ebitenutil.NewImageFromFile("/home/agniva/Downloads/gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		img: img,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
