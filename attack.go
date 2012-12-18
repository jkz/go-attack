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

type Config struct {
	tickrate, hangticks, fallticks, swapticks, clearticks int
}

type Game struct {
	width, height, colors int
	combo, chain          int
	rows                  [][]Block
	config                *Config
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
	x, y                      int
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
	return b.color == AIR && b.counter == 0
}

func (b *Block) IsSupport() bool {
	return b.state != FALL
}

/*
This tick phase checks all blocks for their individual state.
*/
func (b *Block) Tick(under *Block) {
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

/* Create a new rows array and fill it with the old shifted 1 up */
func (g *Game) Push() {
	rows := NewRows(g.width, g.height+1)
	for y, row := range g.rows {
		for x, block := range row {
			fmt.Println(x, y)
			fmt.Println(rows[x][y])
			rows[x][y+1] = block
		}
	}
}

func (g *Game) MoveTick(x int) {
	var y int
	for y = 1; y < g.height; y++ {
		g.rows[y][x].Tick(&g.rows[y-1][x])
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

/* Create a new array of rows with one extra for spawning blocks */
func NewRows(width, height int) [][]Block {
	rows := make([][]Block, height, 1+height)

	for i, _ := range rows {
		rows[i] = make([]Block, width)
	}
	return rows
}

/* NewGame Initializes a game with a viewport of x * y and given amount of
 * block colors */
func NewGame(width, height, colors int) *Game {
	// Set a buffer capacity of twice the height to allow for some garbage
	// and an extra line for spawning blocks
	return &Game{
		width:  width,
		height: height,
		colors: colors,
		rows:   NewRows(width, height)}
}

var FRAME int = 0

func main() {
	grid := colorgrid.Grid{Cell: colorgrid.Size{5, 3}}
	player := newPlayer()
	out := make(chan keydown.Op)
	go keydown.Enable(&cursor, out)

	ch := make(chan int, 10)
	//cX := cursor.X
	//cY := game.height - cursor.Y
	go runTicker(ch)
	//gch := make(chan *Game, 10)
	//go render(gch)
	for i := range ch {
		/*
			select {
			case <-out:
				cX = cursor.X
				cY = game.height - cursor.Y
			default:
			}
		*/
		cX := cursor.X
		cY := game.height - cursor.Y
		//gch <- &game
		fmt.Println(i)
		render(&game, cX, cY)
		game.Tick()
		FRAME += 1
		game.grid.Render(1, 0, fmt.Sprintf("%d", FRAME), colorgrid.WHITE,
			colorgrid.BLACK)
		//fmt.Printf("%d", FRAME)
	}
}
