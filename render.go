package main

import (
	"fmt"
	"github.com/jessethegame/colorgrid"
)

func (g colorgrid.Grid) Block(b Block) {
	g.Render(x, y, fmt.Sprintf("%d", b.counter), colorgrid.BLACK, b.color)
}

func (g colorgrid.Grid) Cursor(c Cursor) {
	g.Render(c.x, c.y, "[", colorgrid.BLACK, colorgrid.WHITE)
	g.Render(c.x+1, c.y, "]", colorgrid.BLACK, colorgrid.WHITE)
}

func (g colorgrid.Grid) Game(game Game) {
	for y, row := range game.rows {
		for x, block := range row {
			g.Block(block)
		}
	}
	g.Cursor(game.cursor)
}

func render(game *Game, cX, cY int) {
	//fmt.Printf("\x1b[H")
	for y, row := range game.rows {
		for x, block := range row {
			if x == cX && y == cY {
				game.grid.Render(cX, cY, "[", colorgrid.BLACK, colorgrid.WHITE)
			} else if x == cX+1 && y == cY {
				game.grid.Render(cX+1, cY, "]", colorgrid.BLACK, colorgrid.WHITE)
			} else {
				game.grid.Render(game.width-x, game.height-y, fmt.Sprintf("%d",
					block.counter), colorgrid.BLACK, block.color)
			}
		}
	}
}
