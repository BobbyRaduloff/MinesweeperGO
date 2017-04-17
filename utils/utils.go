package utils

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

func ImageOrDie(path string) *ebiten.Image {
	i, _, e := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func GetMouseState() (bool, bool, int, int) {
	x, y := ebiten.CursorPosition()
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft), ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight), x, y
}
