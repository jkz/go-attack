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

// ticks / second
var tickrate = 1

// frames / tick
var framerate = 1

// Time in frames
var hangtime uint = 10
var falltime uint = 4
var swaptime uint = 4
var cleartime uint = 4

type State int
type Color int
type ColorCode string

/*
const (
	WHITE Color = "\033[01;37m";
	RED = "\033[22;31m";
	GREEN = "\033[22;32m";
	LIGHT_BLUE = "\033[01;34m";
	MAGENTA = "\033[22;35m";
	YELLOW = "\033[01;33m";
*/
const (
	AIR Color = iota
	ROCK
	// Everything > ROCK is a color
	WHITE
	RED
	GREEN
	LIGHT_BLUE
	MAGENTA
	YELLOW
	BLACK
)

var COLORS = map[Color]ColorCode{
	WHITE:      "\033[01;47m",
	RED:        "\033[22;41m",
	GREEN:      "\033[22;42m",
	LIGHT_BLUE: "\033[01;44m",
	MAGENTA:    "\033[22;45m",
	YELLOW:     "\033[01;43m",
	BLACK:      "\033[22;40m",
}

const (
	STATIC State = iota
	HANG
	FALL
	SWAP
)

type Control struct {
	Up, Down, Left, Right, Swap, Boost, Pause int
}

type Cursor struct {
	Y        uint
	X        int
	Game     *Game
	Controls *Control
}

/*
func (c *Cursor) Swap(x1, y1, x2, y2 int) {
    g = c.game.rows
    g[y1][x2], g[y2][x2] = g[y2][x2], g[y1][x1]
    g[c.y][c.x].Morph(g[c.y][c.x + 1], swaptime)
    g[c.y][c.x + 1].Morph(g[c.y][c.x], swaptime)
}

/*
func (c *Cursor) Swap() {
    g = c.game.rows
    if g[c.y][c.x].IsSwappable() && g[c.y][c.x + 1].IsSwappable() {
        g[c.y][c.x], g[c.y][c.x + 1] = g[c.y][c.x + 1], g[c.y][c.x]
        g[c.y][c.x].Morph(g[c.y][c.x + 1], swaptime)
        g[c.y][c.x + 1].Morph(g[c.y][c.x], swaptime)
    }
}
*/

type Garbage struct {
}

type Clear struct {
}

type Block struct {
	color   Color
	state   State
	counter uint
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
	width, height, colors uint
	rows                  [][]Block
	cursor                Cursor
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

func (g *Game) MoveTick(x uint) {
	var y uint
	for y = 1; y < g.height; y++ {
		g.rows[y][x].Tick(&g.rows[y-1][x])
	}
}

//func (g *Game)

func (g *Game) Tick() {
	// decr timers
	// check movement
	// set timers
	var x uint
	for x = 0; x < g.width; x++ {
		//go g.MoveTick(x)
		go g.MoveTick(x)
	}
	// check clear
	// spawn garbage
}

/* Create a new array of rows with one extra for spawning blocks */
func NewRows(width, height uint) [][]Block {
	rows := make([][]Block, height, 1+height)

	for i, _ := range rows {
		rows[i] = make([]Block, width)
	}
	return rows
}

/* NewGame Initializes a game with a viewport of x * y and given amount of
 * block colors */
func NewGame(width, height, colors uint) Game {
	// Set a buffer capacity of twice the height to allow for some garbage
	// and an extra line for spawning blocks
	return Game{
		width:  width,
		height: height,
		colors: colors,
		rows:   NewRows(width, height)}
}

func printf(text string, color Color) {
	fmt.Printf(string(COLORS[color]))
	fmt.Printf(text)
	fmt.Printf(string(COLORS[BLACK]))
}

var FRAME int = 0

func render(game *Game) {
	fmt.Printf("\x1b[H")
	var y uint
	for y = game.height - 1; y > 0; y-- {
		for i := 0; i < 3; i++ {
			for x, block := range game.rows[y] {
				var str string
				switch {
				case i == 1:
					fmt.Printf("    %d ", block.counter)
				case game.cursor.Y == y && game.cursor.X == x:
					str = ">     "
				case game.cursor.Y == y && game.cursor.X == x-1:
					str = "     <"
				default:
					str = "      "
				}
				printf(str, block.color)
			}
			printf("                   ", BLACK)
			fmt.Println("")
		}
		//fmt.Println("")
	}
}

/*
var interval = time.Second / 10

func runTicker(ch chan int) {
	for _ = range time.Tick(time.Second / 10) {
		ch <- 1
	}
}
*/

func main() {
	game := NewGame(6, 12, 5)
	game.cursor = Cursor{X: 3, Y: 5}
	game.rows[0][0] = Block{color: RED}
	game.rows[0][1] = Block{color: LIGHT_BLUE}
	game.rows[0][2] = Block{color: GREEN}
	game.rows[0][3] = Block{color: MAGENTA}
	game.rows[0][4] = Block{color: RED}
	game.rows[0][5] = Block{color: YELLOW}
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
	game.rows[10][0] = Block{color: RED}
	game.rows[10][1] = Block{color: LIGHT_BLUE}
	game.rows[10][2] = Block{color: GREEN}
	game.rows[10][3] = Block{color: MAGENTA}
	game.rows[10][4] = Block{color: RED}
	game.rows[10][5] = Block{color: YELLOW}
	game.rows[11][3] = Block{color: RED}
	game.rows[11][4] = Block{color: LIGHT_BLUE}
	game.rows[11][5] = Block{color: GREEN}
	game.rows[11][0] = Block{color: MAGENTA}
	game.rows[11][1] = Block{color: RED}
	game.rows[11][2] = Block{color: YELLOW}

	ch := make(chan int, 10)
	go runTicker(ch)
	for i := range ch {
		render(&game)
		fmt.Println(i)
		game.Tick()
		FRAME += 1
		fmt.Printf("%d", FRAME)
	}

}
