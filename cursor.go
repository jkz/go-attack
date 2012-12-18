package main

import (
	"fmt"
	"github.com/jessethegame/keydown"
	"github.com/nsf/termbox-go"
)

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

type coords struct {
	x, y int
}

/* Cursor encloses positions (x, y) and (x + 1, y), so max x coordinate
 * game.width - 2 */
type Cursor struct {
	keydown.Controller
	X, Y int
	Game *Game
}

func (c *Cursor) Up() (int, int) {
	if c.Y < c.Game.height-2 {
		c.Y++
	}
}

func (c *Cursor) Down() {
	if c.Y > 1 {
		c.Y--
	}
}

func (c *Cursor) Right() {
	if c.X < c.Game.width-2 {
		c.X++
	}
}

func (c *Cursor) Left() {
	if c.X > 0 {
		c.X--
	}
}

func (c *Cursor) KeyDown(in chan termbox.Key, out chan keydown.Op) {
	for {
		select {
		case key := <-in:
			switch key {
			case termbox.KeyArrowUp:
				c.Up()
			case termbox.KeyArrowDown:
				c.Down()
			case termbox.KeyArrowLeft:
				c.Left()
			case termbox.KeyArrowRight:
				c.Right()
			case ' ':
				fmt.Print("SWAP")
			case 'c':
				fmt.Print("PUSH")
			case termbox.KeyEsc:
				out <- true
				return
			default:
				panic("Illegal operation!")
			}
			out <- true
		}
	}
}
