package main

import (
	"github.com/jessethegame/colorgrid"
)

func (b Block) render(g colorgrid.Grid) {
	g.Cell(b.pos.x, b.pos.y, rune(b.counter+'0'), colorgrid.WHITE, b.color)
}

func (b Block) cursor(g colorgrid.Grid) {
	g.Cell(b.pos.x, b.pos.y, '[', colorgrid.BLACK, colorgrid.WHITE)
	g.Cell(b.pos.x+1, b.pos.y, ']', colorgrid.BLACK, colorgrid.WHITE)
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
	g.Clear()
	g.Flush()
	p.game.render(g)
	p.cursor.cursor(g)
	//g.Cursor(p.Block)
}
