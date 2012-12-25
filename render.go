package main

import (
	"fmt"
	"github.com/jessethegame/colorgrid"
)

func (b Block) render(g colorgrid.Grid) {
	g.Render(b.pos.x, b.pos.y, fmt.Sprintf("%d", b.counter), colorgrid.BLACK, b.color)
}

func (b Block) cursor(g colorgrid.Grid) {
	g.Render(b.pos.x, b.pos.y, "[", colorgrid.BLACK, colorgrid.WHITE)
	g.Render(b.pos.x+1, b.pos.y, "]", colorgrid.BLACK, colorgrid.WHITE)
}

func (game Game) render(g colorgrid.Grid) {
	for x, col := range game.blocks {
		for y, block := range col {
			block.pos.x = x
			block.pos.y = y
			block.render(g)
		}
	}
}

func (p Player) render(g colorgrid.Grid) {
	p.game.render(g)
	p.cursor.cursor(g)
	//g.Cursor(p.Block)
}
