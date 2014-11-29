package main

import (
// "fmt"
)

func (b *Block) IsSwappable() bool {
	return b.counter == 0
}

func (b *Block) IsEmpty() bool {
	return b.counter == 0 && b.color == AIR
}

func (b *Block) IsSupport() bool {
	return b.state != FALL
}

func (b *Block) IsClearable() bool {
	return b.IsSwappable() && b.under != nil && b.under.IsSupport() && b.color != AIR
}

func (b *Block) IsComboable() bool {
	return b.IsClearable() || (b.state == CLEAR && b.counter == clearticks)
}

func (b *Block) Clear() int {
	if b.state == CLEAR {
		return 0
	}
	b.counter = clearticks
	b.state = CLEAR
	b.chain = true
	return 1
}

func (b *Block) Combo() int {
	var combo = 0

	if b.left.IsComboable() && b.right.IsComboable() {
		if b.left.color == b.color && b.right.color == b.color {
			combo += b.Clear()
			combo += b.left.Clear()
			combo += b.right.Clear()
		}
	}

	if b.under.IsComboable() && b.above.IsComboable() {
		if b.under.color == b.color && b.above.color == b.color {
			combo += b.Clear()
			combo += b.under.Clear()
			combo += b.above.Clear()
		}
	}

	return combo
}

func (b *Block) UpdateState() {
	// If the block has a counter, decrement it, return if it is not done
	if b.counter > 0 {
		b.counter--
		if b.counter > 0 {
			return
		}
	}

	switch b.state {
	case STATIC, SWAP:
		switch {
		case b.color == AIR:
			return
		case b.under == Wall:
			b.state = STATIC
		case b.under.state == HANG:
			b.state = HANG
			b.counter = b.under.counter
			b.chain = b.under.chain
		case b.under.IsEmpty():
			b.state = HANG
			b.counter = hangticks
		}
	case HANG:
		b.state = FALL
		fallthrough
	case FALL:
		switch {
		case b.under.IsEmpty():
			b.under.Copy(b)
			b.Erase()
		case b.under.state == CLEAR:
			b.state = STATIC
		default:
			b.state = b.under.state
			b.counter = b.under.counter
		}
	case CLEAR:
		b.Erase()
		b.counter = hangticks
	default:
		panic("Unknown state!")
	}

}

func (g *Game) UpdateNeighbors() {
	var block *Block
	for x, col := range g.blocks {
		for y, _ := range col {
			block = &g.blocks[x][y]

			if x > 0 {
				block.left = &g.blocks[x-1][y]
			} else {
				block.left = Wall
			}

			if x < g.width-1 {
				block.right = &g.blocks[x+1][y]
			} else {
				block.right = Wall
			}

			if y > 0 {
				block.under = &g.blocks[x][y-1]
			} else {
				block.under = Wall
			}

			if y < g.height-1 {
				block.above = &g.blocks[x][y+1]
			} else {
				block.above = Wall
			}

		}
	}
}

func (g *Game) UpdateState() {
	for x, col := range g.blocks {
		for y, _ := range col {
			g.blocks[x][y].UpdateState()
		}
	}
}

func (g *Game) UpdateCombo() int {
	var combo = 0

	for x, col := range g.blocks {
		for y, _ := range col {
			combo += g.blocks[x][y].Combo()
		}
	}

	return combo
}

func (g *Game) Swap(c coord) {
	if !g.blocks[c.x][c.y].IsSwappable() || !g.blocks[c.x+1][c.y].IsSwappable() {
		return
	}
	g.blocks[c.x][c.y], g.blocks[c.x+1][c.y] = g.blocks[c.x+1][c.y], g.blocks[c.x][c.y]
}

func (g *Game) Tick() {
	g.UpdateNeighbors()
	g.UpdateState()
	var combo = g.UpdateCombo()
	// TODO this is incorrect at the moment
	if combo > 0 {
		g.chain++
	} else {
		g.chain = 1
	}
	// spawn garbage
}
