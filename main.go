package main

import (
	"github.com/hajimehoshi/ebiten"
	"minesweeper/minefield"
	"minesweeper/utils"
	"os"
)

const (
	screenWidth  = minesWidth * 16
	screenHeight = minesHeight * 16
	minesHeight  = 24
	minesWidth   = 34
	minesN       = 90
)

var (
	field   [][]int
	sprites *ebiten.Image
)

type imageParts struct {
	sprite int
	x      int
	y      int
}

func (p *imageParts) Src(i int) (int, int, int, int) {
	y := 0
	x := p.sprite * 16
	if p.sprite > 8 {
		y = 16
		x -= p.sprite * 16 * 8
	}

	return x, y, x + 16, y + 16
}

func (p *imageParts) Dst(i int) (int, int, int, int) {
	return p.y * 16, p.x * 16, p.y*16 + 16, p.x*16 + 16
}

func (p *imageParts) Len() int {
	return 24
}

func update(screen *ebiten.Image) error {
	options := &ebiten.DrawImageOptions{}

	left, _, mx, my := utils.GetMouseState()
	mx /= 16
	my /= 16
	if left  && mx >= 0 && mx < minesWidth && my >= 0 && my < minesHeight {
		switch field[my][mx] {
			case minefield.BOMB:
				field[my][mx] /= minefield.MAGIC
				os.Exit(1)
			case minefield.FLAG:
			case minefield.UNCLICKEDBOMB:
			case minefield.EMPTY:
			default:
				if field[my][mx] % minefield.MAGIC == 0 {
					field[my][mx] /= minefield.MAGIC
				}
		}
	}

	for i := 0; i < minesHeight; i++ {
		for j := 0; j < minesWidth; j++ {
			var bombType int
			switch field[i][j] {
			case minefield.UNCLICKEDBOMB:
				bombType = 8
			case minefield.BOMB / minefield.MAGIC:
				bombType = 9
			case minefield.BOMB:
				bombType = 12
			case minefield.FLAG:
				bombType = 10
			case minefield.CLICKED:
				bombType = 11
			case minefield.EMPTY:
				bombType = 12
			default:
				if field[i][j]%minefield.MAGIC == 0 {
					bombType = 12
				} else {
					bombType = field[i][j] - 1
				}
			}
			options.ImageParts = &imageParts{bombType, i, j}
			screen.DrawImage(sprites, options)
		}
	}
	return nil
}

func init() {
	field = minefield.GenerateField(minesHeight, minesWidth, minesN)
	sprites = utils.ImageOrDie("sprites/sprites.png")
}

func main() {
	ebiten.Run(update, screenWidth, screenHeight, 1, "MinesweeperGO")
}
