package main

import (
	"fmt"
	"github.com/jessethegame/colorgrid"
)

func (b Block) render(g colorgrid.Grid) {
	g.Render(x, y, fmt.Sprintf("%d", b.counter), colorgrid.BLACK, b.color)
}

func (b Block) cursor(g colorgrid.Grid) {
	g.Render(b.x, b.y, "[", colorgrid.BLACK, colorgrid.WHITE)
	g.Render(b.x+1, b.y, "]", colorgrid.BLACK, colorgrid.WHITE)
}

func (game Game) render(g colorgrid.Grid) {
	for y, row := range game.rows {
		for x, block := range row {
			block.render(g)
		}
	}
}

func (p Player) render(g colorgrid.Grid) {
	p.game.render(g)
	g.Cursor(p.Block)
}
