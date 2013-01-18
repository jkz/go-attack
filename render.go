package main

import (
	"fmt"
	"github.com/jessethegame/colorgrid"
)

func (b Block) render(g colorgrid.Grid) {
	//	g.Cell(b.pos.x, b.pos.y, rune(b.counter+'0'), colorgrid.WHITE, b.color)
	g.Cell(b.pos.x, b.pos.y, ' ', colorgrid.WHITE, b.color)
}

func (b Block) cursor(g colorgrid.Grid) {
	fmt.Print("ULUBULU")
	g.Cell(b.pos.x, b.pos.y, '[', colorgrid.BLACK, colorgrid.WHITE)
	g.Cell(b.pos.x+1, b.pos.y, ']', colorgrid.BLACK, colorgrid.WHITE)
}

func (game Game) render(g colorgrid.Grid) {
	for x, col := range game.blocks {
		for y, block := range col {
			block.pos.x = x
			block.pos.y = game.height - y
			block.render(g)
		}
	}
}

func (p Player) render(g colorgrid.Grid) {
	g.Clear()
	//p.game.render(g)
	//p.cursor.cursor(g)
	for x, col := range p.game.blocks {
		for y, block := range col {
			block.pos.x = x
			block.pos.y = p.game.height - y
			block.render(g)
		}
	}
	g.Cell(p.cursor.pos.x, p.game.height-p.cursor.pos.y, '[', colorgrid.WHITE, p.cursor.color)
	g.Cell(p.cursor.pos.x+1, p.game.height-p.cursor.pos.y, ']', colorgrid.WHITE, p.cursor.color)
	g.Flush()
}
