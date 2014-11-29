package main

/*
- Not visible; holds garbage
= visible game area
+ spawns new blocks

y+1  - - - - - -
y    - - - - - -
y-1  = = = = = =
y-2  = = = = = =
:    : : : : : :
2    = = = = = =
1    = = = = = =
0    + + + + + +
*/

import (
	"fmt"
	"github.com/jessethegame/go-colorgrid"
	"github.com/jessethegame/go-keydown"
)

// ticks / second
var tickrate = 1

// Time in frames
var hangticks int = 6
var fallticks int = 4
var swapticks int = 4
var clearticks int = 12

var Wall *Block

type command int

const (
	PUSH command = iota
	PAUSE
	QUIT
)

type Config struct {
	tickrate, hangticks, fallticks, swapticks, clearticks int
}

type coord struct {
	x, y int
}

type Game struct {
	width, height, colors int
	combo, chain          int
	blocks                [][]Block
	config                *Config
	swap                  chan coord
	command               chan command
	paused                bool
}

type State int

const (
	AIR  colorgrid.Color = colorgrid.DEFAULT
	WALL                 = colorgrid.WHITE
	ROCK                 = colorgrid.WHITE
	// Everything > ROCK is a color
)

const (
	STATIC State = iota
	HANG
	FALL
	SWAP
	CLEAR
)

type Garbage struct {
}

type Block struct {
	above, under, left, right *Block
	pos                       coord
	color                     colorgrid.Color
	state                     State
	counter                   int
	chain                     bool
	is_cursor                 bool
	//garbage *Garbage
}

type Cursor struct {
	pos         coord
	color       colorgrid.Color
	left, right *Block
}

func (b *Block) Become(other *Block) {
	b.under, b.above, b.left, b.right, other.under, other.above, other.left,
		other.right = other.under, other.above, other.left, other.right, b.under,
		b.above, b.left, b.right
	if b.under != nil {
		b.under.above = other
	}
	if b.above != nil {
		b.above.under = other
	}
	if b.right != nil {
		b.right.left = other
	}
	if b.left != nil {
		b.left.right = other
	}
	b.color = other.color
	b.state = other.state
	b.counter = other.counter
	//b.garbage = other.garbage
}

func (b *Block) Copy(other *Block) {
	b.state = other.state
	b.color = other.color
	b.counter = other.counter
	b.chain = other.chain
}

func (b *Block) Erase() {
	b.state = STATIC
	b.color = AIR
	b.counter = 0
	b.chain = false
}

/* Create a new blocks array and fill it with the old shifted 1 up */
func (g *Game) Push(height int) {
	blocks := newBlocks(g.width, height)
	for x, col := range g.blocks {
		for y, block := range col {
			fmt.Println(x, y)
			fmt.Println(blocks[x][y])
			blocks[x][y+1] = block
		}
	}
}

/* Create a new array of blocks with one extra for spawning blocks */
func newBlocks(width, height int) [][]Block {
	blocks := make([][]Block, width)

	for i, _ := range blocks {
		// blocks[i] = make([]Block, height+1)
		blocks[i] = make([]Block, height)
	}
	return blocks
}

/* NewGame Initializes a game with a viewport of x * y and given amount of
 * block colors */
func newGame(width, height, colors int) *Game {
	// Set a buffer capacity of twice the height to allow for some garbage
	// and an extra line for spawning blocks
	return &Game{
		width:   width,
		height:  height,
		colors:  colors,
		blocks:  newBlocks(width, height),
		swap:    make(chan coord, 10),
		command: make(chan command, 10)}
}

func (p *Player) Play() {
	grid := colorgrid.NewGrid(5, 3, colorgrid.WHITE, colorgrid.BLACK)
	// grid := colorgrid.NewGrid(2, 1, colorgrid.WHITE, colorgrid.BLACK)
	control := keydown.NewController()
	frames := make(chan int, 10)

	go control.Run()
	go p.listen(control.Input)
	go runTicker(frames)

	for frame := range frames {
		//fmt.Print("F", frame)
		// Yuck, needs reference for frame
		frame++
		select {
		case pos, ok := <-p.game.swap:
			if ok {
				// fmt.Print("SWAP", pos, "\r\n")
				p.game.Swap(p.cursor.pos)
			} else {
				fmt.Print("NOT SWAP", pos, "\r\n")
				control.Stop <- true
			}
			// grid.Print(0, (player.game.height+1)*grid.Size.Height, ".", colorgrid.WHITE, colorgrid.BLACK)
			// fmt.Print(pos)
		case com, ok := <-p.game.command:
			if !ok {
				// fmt.Print("NOT com:", com, "\r\n")
				control.Stop <- true
				return
			}
			fmt.Print("com:", com)
			switch com {
			case PUSH:
				// fmt.Println("PUSH")
			case PAUSE:
				// fmt.Println("PAUSE")
				p.game.paused = !p.game.paused
			case QUIT:
				// fmt.Println("QUIT")
				control.Stop <- true
				return
			default:
				// fmt.Print("defaulted")
			}
		default:
		}
		if !p.game.paused {
			p.game.Tick()
			p.render(grid)
			// player.game.render(grid)
		}
	}
}

func main() {
	Wall = &Block{
		color: colorgrid.WHITE}

	player := defaultPlayer()
	player.Play()
}
