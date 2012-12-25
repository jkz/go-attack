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

import "fmt"
import "github.com/jessethegame/keydown"
import "github.com/jessethegame/colorgrid"

// ticks / second
var tickrate = 1

// Time in frames
var hangticks int = 10
var fallticks int = 4
var swapticks int = 4
var clearticks int = 4

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
}

type State int

const (
	AIR  colorgrid.Color = colorgrid.BLACK
	ROCK                 = colorgrid.GRAY
	// Everything > ROCK is a color
)

const (
	STATIC State = iota
	HANG
	FALL
	SWAP
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
	garbage                   *Garbage
}

func (b *Block) Become(other *Block) {
	b.color = other.color
	b.state = other.state
	b.counter = other.counter
	b.garbage = other.garbage
}

func (b *Block) IsSwappable() bool {
	return b.counter == 0
}

func (b *Block) IsEmpty() bool {
	return b.counter == 0 && b.color == AIR
}

func (b *Block) IsSupport() bool {
	return b.state != FALL
}

/*
This tick phase checks all blocks for their individual state.
*/
func (b *Block) Tick(under *Block) {
	// If the block has a counter, decrement it, return if it is not done
	if b.counter > 0 {
		b.counter--
		if b.counter > 0 {
			return
		}
	}

	switch b.state {
	case STATIC, SWAP:
		if b.color == AIR {
			return
		}
		switch {
		case under.state == HANG:
			b.state = HANG
			b.counter = under.counter
			b.chain = under.chain
		case under.IsEmpty():
			b.state = HANG
			b.counter = hangticks
		}
	case HANG:
		b.state = FALL
		fallthrough
	case FALL:
		if under.IsEmpty() {
			under.Become(b)
			b.Become(&Block{})
		} else {
			b.state = under.state
			b.counter = under.counter
		}
	default:
		panic("I'm not supposed to get here!")
	}
}

/* Create a new blocks array and fill it with the old shifted 1 up */
func (g *Game) Push() {
	blocks := newBlocks(g.width, g.height+1)
	for y, row := range g.blocks {
		for x, block := range row {
			fmt.Println(x, y)
			fmt.Println(blocks[x][y])
			blocks[x][y+1] = block
		}
	}
}

func (g *Game) MoveTick(x int) {
	var y int
	for y = 1; y < g.height; y++ {
		g.blocks[y][x].Tick(&g.blocks[y-1][x])
	}
}

func (g *Game) Tick() {
	// decr timers
	// check movement
	// set timers
	var x int
	for x = 0; x < g.width; x++ {
		//go g.MoveTick(x)
		go g.MoveTick(x)
	}
	// check clear
	// spawn garbage
}

/* Create a new array of blocks with one extra for spawning blocks */
func newBlocks(width, height int) [][]Block {
	blocks := make([][]Block, width)

	for i, _ := range blocks {
		blocks[i] = make([]Block, height+1)
	}
	return blocks
}

/* NewGame Initializes a game with a viewport of x * y and given amount of
 * block colors */
func newGame(width, height, colors int) *Game {
	// Set a buffer capacity of twice the height to allow for some garbage
	// and an extra line for spawning blocks
	return &Game{
		width:  width,
		height: height,
		colors: colors,
		blocks: newBlocks(width, height)}
}

func main() {
	//grid := colorgrid.Grid{Cell: colorgrid.Size{5, 3}}
	fmt.Println("START")
	control := keydown.NewController()
	player := defaultPlayer()
	fmt.Println("INPUT")
	frames := newTicker()
	go control.Run()
	go player.listen(control.Input)
	go runTicker(frames)
	for frame := range frames {
		fmt.Print("FRAME", frame, "\n")
		select {
		case pos := <-player.game.swap:
			fmt.Printf("SWAP", pos, "\n")
		case com := <-player.game.command:
			switch com {
			case PUSH:
				fmt.Println("PUSH")
			case PAUSE:
				fmt.Println("PAUSE")
			case QUIT:
				fmt.Println("QUIT")
				control.Stop <- true
			}
		default:
			fmt.Print(".")
		}
	}
}
