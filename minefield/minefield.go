package minefield

import (
	"math/rand"
	"time"
)

const (
	UNCLICKEDBOMB = 0xFAAF
	BOMB          = 0xDEAD
	FLAG          = 0xBEEF
	CLICKED       = 0xF00F
	EMPTY         = 0
	MAGIC         = 41
)

func GenerateField(height int, width int, bombs int) [][]int {
	field := make([][]int, height)
	for i := 0; i < height; i++ {
		field[i] = make([]int, width)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < bombs; {
		x := rand.Intn(height - 1)
		y := rand.Intn(width - 1)
		if field[x][y] == BOMB {
			continue
		}
		field[x][y] = BOMB
		i++
		if x > 1 {
			if field[x-1][y] != BOMB {
				field[x-1][y] += 1
			}
			if y > 1 && field[x-1][y-1] != BOMB {
				field[x-1][y-1] += 1
			}
			if y < width-1 && field[x-1][y+1] != BOMB {
				field[x-1][y+1] += 1
			}
		}
		if x < height-1 {
			if field[x+1][y] != BOMB {
				field[x+1][y] += 1
			}
			if y > 1 && field[x+1][y-1] != BOMB {
				field[x+1][y-1] += 1
			}
			if y < width-1 && field[x+1][y+1] != BOMB {
				field[x+1][y+1] += 1
			}
		}
		if y > 1 && field[x][y-1] != BOMB {
			field[x][y-1] += 1
		}
		if y < width-1 && field[x][y+1] != BOMB {
			field[x][y+1] += 1
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] != BOMB && field[i][j] != FLAG && field[i][j] != EMPTY && field[i][j] != CLICKED && field[i][j] != UNCLICKEDBOMB {
				field[i][j] *= MAGIC
			}
		}
	}
	return field
}
