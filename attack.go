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
	ROCK                 = colorgrid.WHITE
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
	//above, under, left, right *Block
	pos     coord
	color   colorgrid.Color
	state   State
	counter int
	//chain   bool
	//garbage *Garbage
}

func (b *Block) Become(other *Block) {
	b.color = other.color
	b.state = other.state
	b.counter = other.counter
	//b.garbage = other.garbage
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
	fmt.Print("Tick", under, "\r\n")

	switch b.state {
	case STATIC, SWAP:
		fmt.Print("STATIC, SWAP")
		if b.color == AIR {
			fmt.Print("-AIR", b.color)
			return
		}
		switch {
		case under.state == HANG:
			fmt.Print("-HANG")
			b.state = HANG
			b.counter = under.counter
			b.chain = under.chain
		case under.IsEmpty():
			fmt.Print("-EMPTY")
			b.state = HANG
			b.counter = hangticks
		default:
			fmt.Print("-DEFAULT")
		}
	case HANG:
		fmt.Print("HANG")
		b.state = FALL
		fallthrough
	case FALL:
		fmt.Print("FALL")
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
	for x, col := range g.blocks {
		for y, block := range col {
			fmt.Println(x, y)
			fmt.Println(blocks[x][y])
			blocks[x][y+1] = block
		}
	}
}

func (g *Game) MoveTick(x int) {
	for y := 1; y < g.height; y++ {
		g.blocks[x][y].Tick(&g.blocks[x][y-1])
	}
}

func (g *Game) Tick() {
	// decr timers
	// check movement
	// set timers
	var x int
	for x = 0; x < g.width; x++ {
		//go g.MoveTick(x)
		g.MoveTick(x)
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
		width:   width,
		height:  height,
		colors:  colors,
		blocks:  newBlocks(width, height),
		swap:    make(chan coord, 10),
		command: make(chan command, 10)}
}

func main() {
	grid := colorgrid.NewGrid(5, 3, colorgrid.WHITE, colorgrid.BLACK)
	fmt.Println("START")
	control := keydown.NewController()
	player := defaultPlayer()
	fmt.Println("INPUT")
	frames := make(chan int, 10)
	go control.Run()
	go player.listen(control.Input)
	go runTicker(frames)
	for frame := range frames {
		//fmt.Print("F", frame)
		frame++
		select {
		case pos, ok := <-player.game.swap:
			if ok {
				fmt.Print("SWAP", pos, "\r\n")
			} else {
				fmt.Print("NOT SWAP", pos, "\r\n")
				control.Stop <- true
			}
		case com, ok := <-player.game.command:
			if !ok {
				fmt.Print("NOT com:", com, "\r\n")
				control.Stop <- true
				return
			}
			fmt.Print("com:", com)
			switch com {
			case PUSH:
				fmt.Println("PUSH")
			case PAUSE:
				fmt.Println("PAUSE")
			case QUIT:
				fmt.Println("QUIT")
				control.Stop <- true
				return
			default:
				fmt.Print("defaulted")
			}
		default:
			fmt.Print(".")
		}
		fmt.Print("TICK")
		player.game.Tick()
		fmt.Print("RENDER")
		player.game.render(grid)
	}
}
