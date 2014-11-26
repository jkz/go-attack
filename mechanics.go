package main

import (
	"fmt"
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
	fmt.Print(b.under)
	return b.IsSwappable() && b.under != nil && b.under.IsSupport()
}

func (b *Block) IsComboable() bool {
	return b.IsClearable() || (b.state == CLEAR && b.counter == clearticks)
}

func (b *Block) Clear() int {
	if b.IsClearable() {
		return b.Combo()
	} else {
		return 0
	}
}

// This is a bit weird, it will clear combos per three horizontally
// and vertically. To keep conditions simple, count the number of blocks
// in the combo and propagate it, this function recursively performs
// the same check on the neighbours. This means that in one combo, a single
// block might be checked 6 times (2 x 3)
func (b *Block) Combo() int {
	var combo = 0

	if b.counter > 0 {
		return 0
	}

	if b.left.IsComboable() && b.right.IsComboable() {
		if b.left.color == b.color && b.right.color == b.color {
			combo += b.left.Combo()
			combo += b.right.Combo()
		}
	}

	if b.under.IsComboable() && b.above.IsComboable() {
		if b.under.color == b.color && b.above.color == b.color {
			combo += b.under.Combo()
			combo += b.above.Combo()
		}
	}

	if combo > 0 {
		b.counter = clearticks
		b.state = CLEAR
		b.chain = true
		combo++
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
		case b.under == nil:
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
		if b.under.IsEmpty() {
			b.under.Become(b)
			b.Become(&Block{})
		} else {
			b.state = b.under.state
			b.counter = b.under.counter
		}
	default:
		panic("I'm not supposed to get here!")
	}

}

func (g *Game) UpdateRelativePositions() {
	for x, col := range g.blocks {
		for y, block := range col {
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
	for _, col := range g.blocks {
		for _, block := range col {
			block.UpdateState()
		}
	}
}

func (g *Game) UpdateCombo() int {
	var combo = 0

	for _, col := range g.blocks {
		for _, block := range col {
			combo += block.Clear()
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
	g.UpdateRelativePositions()
	g.UpdateState()
	var combo = g.UpdateCombo()
	if combo > 0 {
		g.chain++
	} else {
		g.chain = 1
	}
	// spawn garbage
}
