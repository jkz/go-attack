package main

import (
	"github.com/jessethegame/go-colorgrid"
)

func (b Block) render(g colorgrid.Grid) {
	//	g.Cell(b.pos.x, b.pos.y, rune(b.counter+'0'), colorgrid.WHITE, b.color)
	g.Cell(b.pos.x, b.pos.y, ' ', colorgrid.WHITE, b.color)
}

// Render each Block
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
	// g.Clear()
	// p.game.render(g)
	//p.cursor.cursor(g)

	var sym rune

	for x, col := range p.game.blocks {
		for y, block := range col {
			block.pos.x = x
			block.pos.y = p.game.height - y
			switch {
			case p.cursor.pos.x == block.pos.x && p.cursor.pos.y == y:
				sym = '['
			case p.cursor.pos.x+1 == block.pos.x && p.cursor.pos.y == y:
				sym = ']'
			case block.counter > 0:
				sym = rune('0' + block.counter)
			case block.state == SWAP:
				sym = 'W'
			case block.state == HANG:
				sym = 'H'
			case block.state == FALL:
				sym = 'F'
			case block.state == CLEAR:
				sym = 'C'
			default:
				sym = ' '
			}

			// block.render(g)
			g.Cell(block.pos.x, block.pos.y, sym, colorgrid.WHITE, block.color)
		}
	}

	// Render the cursor
	var left, right *Block
	left = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y]
	right = &p.game.blocks[p.cursor.pos.x+1][p.cursor.pos.y]
	g.Cell(p.cursor.pos.x, p.game.height-p.cursor.pos.y, '[', colorgrid.WHITE, left.color)
	g.Cell(p.cursor.pos.x+1, p.game.height-p.cursor.pos.y, ']', colorgrid.WHITE, right.color)
	// g.Flush()
	g.Cell(0, 0, ' ', colorgrid.WHITE, colorgrid.BLACK)
}
