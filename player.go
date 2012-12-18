package main

import (
	"github.com/nsf/termbox-go"
)

type keys struct {
	Up, Down, Left, Right, Swap, Push, Pause, Quit,
	Top, Bottom, Home, End termbox.Key
}

type Player struct {
	keys   keys
	cursor *Block
	game   *Game
}

/*
A key channel listener. It directly performs any operations on the cursor;
propagates other operations to channels on the game.
*/
func (p *Player) pressed(ch chan termbox.Key) {
	for key := range ch {
		switch key {
		case p.keys.Up:
			if p.cursor.y < p.game.height {
				p.cursor = &p.game.blocks[cursor.x][cursor.y+1]
			}
		case p.keys.Down:
			if p.cursor.y > 0 {
				p.cursor = &p.game.blocks[cursor.x][cursor.y-1]
			}
		case p.keys.Left:
			if p.cursor.x > 0 {
				p.cursor = &p.game.blocks[cursor.x-1][cursor.y]
			}
		case p.keys.Right:
			if p.cursor.x < p.game.width-2 {
				p.cursor = &p.game.blocks[cursor.x+1][cursor.y]
			}
		case p.keys.Top:
			p.cursor = &p.game.blocks[cursor.x][p.game.height-1]
		case p.keys.Bottom:
			p.cursor = &p.game.blocks[cursor.x][0]
		case p.keys.Home:
			p.cursor = &p.game.blocks[0][p.cursor.y]
		case p.keys.End:
			p.cursor = &p.game.blocks[p.game.width-2][p.cursor.y]
		case p.keys.Swap:
			p.game.swap <- p.cursor
		case p.keys.Push:
			p.game.push <- true
		case p.keys.Quit:
			panic("Quit!")
		case p.keys.Pause:
		default:
			panic("Illegal command!")
		}
	}
}
