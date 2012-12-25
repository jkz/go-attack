package main

import (
	"fmt"
	"github.com/jessethegame/colorgrid"
)

type keys struct {
	up, down, left, right, swap, push, pause, quit,
	top, bottom, home, end keydown.Key
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
func (p *Player) listen(ch chan keydown.Key) {
	fmt.Println("LISTENING")
	for {
		select {
		case key := <-ch:
			fmt.Print("KEY", key)
			switch key {
			case p.keys.up:
				fmt.Print("UP")
				if p.cursor.pos.y < p.game.height {
					p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y+1]
				}
			case p.keys.down:
				fmt.Print("DOWN")
				if p.cursor.pos.y > 0 {
					p.cursor = &p.game.blocks[p.cursor.pos.x][p.cursor.pos.y-1]
				}
			case p.keys.left:
				fmt.Print("LEFT")
				if p.cursor.pos.x > 0 {
					p.cursor = &p.game.blocks[p.cursor.pos.x-1][p.cursor.pos.y]
				}
			case p.keys.right:
				fmt.Print("RIGHT")
				if p.cursor.pos.x < p.game.width-2 {
					p.cursor = &p.game.blocks[p.cursor.pos.x+1][p.cursor.pos.y]
				}
			case p.keys.top:
				p.cursor = &p.game.blocks[p.cursor.pos.x][p.game.height-1]
				fmt.Print("TOP")
			case p.keys.bottom:
				p.cursor = &p.game.blocks[p.cursor.pos.x][0]
				fmt.Print("BOT")
			case p.keys.home:
				p.cursor = &p.game.blocks[0][p.cursor.pos.y]
				fmt.Print("HOME")
			case p.keys.end:
				p.cursor = &p.game.blocks[p.game.width-2][p.cursor.pos.y]
				fmt.Print("END")
			case p.keys.swap:
				fmt.Print("SWAP")
				p.game.swap <- coord{p.cursor.pos.x, p.cursor.pos.y}
			case p.keys.push:
				fmt.Print("push?")
				p.game.command <- PUSH
			case p.keys.quit:
				fmt.Print("quit?")
				p.game.command <- QUIT
			case p.keys.pause:
				fmt.Print("pause?")
				p.game.command <- PAUSE
			default:
				fmt.Print("DEFAULTED")
				//panic("Illegal command!")
			}
		}
	}
}
