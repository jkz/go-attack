package main

import (
	// "fmt"
	"github.com/jessethegame/go-keydown"
)

type keys struct {
	up, down, left, right, swap, push, pause, quit,
	top, bottom, home, end keydown.Key
}

type Player struct {
	keys   keys
	cursor Cursor
	game   *Game
}

/*
A key channel listener. It directly performs any operations on the cursor;
propagates other operations to channels on the game.
*/
func (p *Player) listen(ch chan keydown.Key) {
	for {
		select {
		case key := <-ch:
			// fmt.Print("KEY", key)
			// fmt.Print(p.cursor)
			switch key {
			case p.keys.up:
				if p.cursor.pos.y < p.game.height {
					p.cursor.pos.y += 1
					// p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y+1]
				}
				// if p.cursor.above != nil {
				// 	p.cursor = p.cursor.above
				// }
				// fmt.Print("UP", p.cursor.pos, p.cursor.pos.y)
			case p.keys.down:
				if p.cursor.pos.y > 0 {
					p.cursor.pos.y -= 1
					// p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y-1]
				}
				// if p.cursor.under != nil {
				// 	p.cursor = p.cursor.under
				// }
				// fmt.Print("DOWN", p.cursor.pos)
			case p.keys.left:
				if p.cursor.pos.x > 0 {
					p.cursor.pos.x -= 1
					// p.cursor = &p.game.blocks[p.cursor.pos.x-1][p.cursor.pos.y]
				}
				// if p.cursor.left != nil {
				// 	p.cursor = p.cursor.left
				// }
				// fmt.Print("LEFT", p.cursor.pos)
			case p.keys.right:
				if p.cursor.pos.x < p.game.width-2 {
					p.cursor.pos.x += 1
					// p.cursor = &p.game.blocks[p.cursor.pos.x+1][p.cursor.pos.y]
				}
				// if p.cursor.right != nil {
				// 	p.cursor = p.cursor.right
				// }
				// fmt.Print("RIGHT", p.cursor.pos)
			case p.keys.top:
				p.cursor.pos.y = p.game.height - 1
				// p.cursor = &p.game.blocks[p.cursor.pos.x][p.game.height-1]
				// fmt.Print("TOP")
			case p.keys.bottom:
				p.cursor.pos.y = 0
				// p.cursor = &p.game.blocks[p.cursor.pos.x][0]
				// fmt.Print("BOT")
			case p.keys.home:
				p.cursor.pos.x = 0
				// p.cursor = &p.game.blocks[0][p.cursor.pos.y]
				// fmt.Print("HOME")
			case p.keys.end:
				p.cursor.pos.x = p.game.width - 2
				// p.cursor = &p.game.blocks[p.game.width-2][p.cursor.pos.y]
				// fmt.Print("END")
			case p.keys.swap:
				// fmt.Print("SWAP")
				p.game.swap <- coord{p.cursor.pos.x, p.cursor.pos.y}
			case p.keys.push:
				// fmt.Print("push?")
				p.game.command <- PUSH
			case p.keys.quit:
				// fmt.Print("quit?")
				p.game.command <- QUIT
			case p.keys.pause:
				// fmt.Print("pause?")
				p.game.command <- PAUSE
			default:
				// fmt.Print("DEFAULTED")
				//panic("Illegal command!")
			}
		}
	}
}
