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

// frames / tick
var framerate = 1

// Time in frames
var hangtime int = 10
var falltime int = 4
var swaptime int = 4
var cleartime int = 4

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

type Clear struct {
}

type Block struct {
	color   colorgrid.Color
	state   State
	counter int
	combo   bool
	garbage *Garbage
	clear   *Clear
}

func (b *Block) Become(other *Block) {
	b.color = other.color
	b.state = other.state
	b.counter = other.counter
	b.combo = other.combo
	b.garbage = other.garbage
	b.clear = other.clear
}

func (b *Block) IsSwappable() bool {
	return b.counter == 0
}

func (b *Block) IsEmpty() bool {
	return b.color == AIR && b.counter == 0
}

func (b *Block) IsSupport() bool {
	return b.state != FALL
	return b.color != AIR && b.state != FALL
}

func (b *Block) Tick(under *Block) {
	if b.counter > 0 {
		b.counter--
		if b.counter > 0 {
			return
		}
	}
	//fmt.Println(b, under)

	switch b.state {
	case STATIC, SWAP:
		//fmt.Println("SWITCH", b.state)
		if b.color == AIR {
			return
		}
		switch {
		case under.state == HANG:
			b.state = HANG
			b.counter = under.counter
			b.combo = under.combo
		case under.IsEmpty():
			b.state = HANG
			b.counter = hangtime
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

type Game struct {
	width, height, colors int
	rows                  [][]Block
	cursor                Cursor
	grid                  colorgrid.Grid
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

//func (g *Game)

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
func NewGame(width, height, colors int) Game {
	// Set a buffer capacity of twice the height to allow for some garbage
	// and an extra line for spawning blocks
	return Game{
		width:  width,
		height: height,
		colors: colors,
		rows:   NewRows(width, height)}
}

var FRAME int = 0

func main() {
	game := NewGame(6, 12, 5)
	cursor := Cursor{X: 3, Y: 5, Game: &game}
	game.grid.Cell = colorgrid.Size{5, 3}
	game.rows[0][0] = Block{color: colorgrid.RED}
	game.rows[0][1] = Block{color: colorgrid.LIGHT_BLUE}
	game.rows[0][2] = Block{color: colorgrid.GREEN}
	game.rows[0][3] = Block{color: colorgrid.MAGENTA}
	game.rows[0][4] = Block{color: colorgrid.RED}
	game.rows[0][5] = Block{color: colorgrid.YELLOW}
	/*
		game.rows[1][3] = Block{color: RED}
		game.rows[1][4] = Block{color: LIGHT_BLUE}
		game.rows[1][5] = Block{color: GREEN}
		game.rows[1][0] = Block{color: MAGENTA}
		game.rows[1][1] = Block{color: RED}
		game.rows[1][2] = Block{color: YELLOW}
		game.rows[2][0] = Block{color: RED}
		game.rows[2][1] = Block{color: LIGHT_BLUE}
		game.rows[2][2] = Block{color: GREEN}
		game.rows[2][3] = Block{color: MAGENTA}
		game.rows[2][4] = Block{color: RED}
		game.rows[2][5] = Block{color: YELLOW}
		game.rows[3][3] = Block{color: RED}
		game.rows[3][4] = Block{color: LIGHT_BLUE}
		game.rows[3][5] = Block{color: GREEN}
		game.rows[3][0] = Block{color: MAGENTA}
		game.rows[3][1] = Block{color: RED}
		game.rows[3][2] = Block{color: YELLOW}
		game.rows[4][0] = Block{color: RED}
		game.rows[4][1] = Block{color: LIGHT_BLUE}
		game.rows[4][2] = Block{color: GREEN}
		game.rows[4][3] = Block{color: MAGENTA}
		game.rows[4][4] = Block{color: RED}
		game.rows[4][5] = Block{color: YELLOW}
		game.rows[5][3] = Block{color: RED}
		game.rows[5][4] = Block{color: LIGHT_BLUE}
		game.rows[5][5] = Block{color: GREEN}
		game.rows[5][0] = Block{color: MAGENTA}
		game.rows[5][1] = Block{color: RED}
		game.rows[5][2] = Block{color: YELLOW}
		game.rows[6][0] = Block{color: RED}
		game.rows[6][1] = Block{color: LIGHT_BLUE}
		game.rows[6][2] = Block{color: GREEN}
		game.rows[6][3] = Block{color: MAGENTA}
		game.rows[6][4] = Block{color: RED}
		game.rows[6][5] = Block{color: YELLOW}
		game.rows[7][3] = Block{color: RED}
		game.rows[7][4] = Block{color: LIGHT_BLUE}
		game.rows[7][5] = Block{color: GREEN}
		game.rows[7][0] = Block{color: MAGENTA}
		game.rows[7][1] = Block{color: RED}
		game.rows[7][2] = Block{color: YELLOW}
		game.rows[8][0] = Block{color: RED}
		game.rows[8][1] = Block{color: LIGHT_BLUE}
		game.rows[8][2] = Block{color: GREEN}
		game.rows[8][3] = Block{color: MAGENTA}
		game.rows[8][4] = Block{color: RED}
		game.rows[8][5] = Block{color: YELLOW}
		game.rows[9][3] = Block{color: RED}
		game.rows[9][4] = Block{color: LIGHT_BLUE}
		game.rows[9][5] = Block{color: GREEN}
		game.rows[9][0] = Block{color: MAGENTA}
		game.rows[9][1] = Block{color: RED}
		game.rows[9][2] = Block{color: YELLOW}
	*/
	game.rows[10][0] = Block{color: colorgrid.RED}
	game.rows[10][1] = Block{color: colorgrid.LIGHT_BLUE}
	game.rows[10][2] = Block{color: colorgrid.GREEN}
	game.rows[10][3] = Block{color: colorgrid.MAGENTA}
	game.rows[10][4] = Block{color: colorgrid.RED}
	game.rows[10][5] = Block{color: colorgrid.YELLOW}
	game.rows[11][3] = Block{color: colorgrid.RED}
	game.rows[11][4] = Block{color: colorgrid.LIGHT_BLUE}
	game.rows[11][5] = Block{color: colorgrid.GREEN}
	game.rows[11][0] = Block{color: colorgrid.MAGENTA}
	game.rows[11][1] = Block{color: colorgrid.RED}
	game.rows[11][2] = Block{color: colorgrid.YELLOW}

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
