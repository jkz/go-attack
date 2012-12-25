package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type keys struct {
	up, down, left, right, swap, push, pause, quit,
	top, bottom, home, end termbox.Key
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
func (p *Player) listen(ch chan termbox.Key) {
	for key := range ch {
		switch key {
		case p.keys.up:
			fmt.Println("UP")
			if p.cursor.pos.y < p.game.height {
				p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y+1]
			}
		case p.keys.down:
			fmt.Println("DOWN")
			if p.cursor.pos.y > 0 {
				p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y-1]
			}
		case p.keys.left:
			fmt.Println("LEFT")
			if p.cursor.pos.x > 0 {
				p.cursor = &p.game.blocks[p.cursor.pos.x-1][p.cursor.pos.y]
			}
		case p.keys.right:
			fmt.Println("RIGHT")
			if p.cursor.pos.x < p.game.width-2 {
				p.cursor = &p.game.blocks[p.cursor.pos.x+1][p.cursor.pos.y]
			}
		case p.keys.top:
			p.cursor = &p.game.blocks[p.cursor.pos.x][p.game.height-1]
		case p.keys.bottom:
			p.cursor = &p.game.blocks[p.cursor.pos.x][0]
		case p.keys.home:
			p.cursor = &p.game.blocks[0][p.cursor.pos.y]
		case p.keys.end:
			p.cursor = &p.game.blocks[p.game.width-2][p.cursor.pos.y]
		case p.keys.swap:
			p.game.swap <- coord{p.cursor.pos.x, p.cursor.pos.y}
		case p.keys.push:
			p.game.command <- PUSH
		case p.keys.quit:
			p.game.command <- QUIT
		case p.keys.pause:
			p.game.command <- PAUSE
		default:
			panic("Illegal command!")
		}
	}
}
