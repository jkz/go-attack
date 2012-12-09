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
var tickrate = 10

// frames / tick
var framerate = 1

var hangtime = 10
var falltime = 4
var swaptime = 4
var cleartime = 4

type State int
type Form int

const (
	AIR Form = iota
	GARBAGE
        ROCK
        // Everything > ROCK is a color
)

const (
	STATIC State = iota
	HANG
	FALL
	SWAP
	CLEAR

	TO_AIR
	TO_SOLID
	TO_GARBAGE
)

/* Cursor encloses positions (x, y) and (x + 1, y), so max x coordinate
 * game.width - 2 */
type Cursor struct {
    X, Y int
    Game *Game
    Controls *Control
}

func (c *Cursor) Up() {
    if c.y < game.height - 2 {
        c.y++
    }
}

func (c *Cursor) Down() {
    if c.y > 1 {
        c.y--
    }
}

func (c *Cursor) Right() {
    if c.x < game.width - 2 {
        c.x++
    }
}

func (c *Cursor) Left() {
    if c.x > 0 {
        c.x--
    }
}

func (c *Cursor) Swap() {
    g = c.game.rows
    if g[c.y][c.x].IsSwappable() && g[c.y][c.x + 1].IsSwappable() {
        g[c.y][c.x], g[c.y][c.x + 1] = g[c.y][c.x + 1], g[c.y][c.x]
        g[c.y][c.x].Morph(g[c.y][c.x + 1], swaptime)
        g[c.y][c.x + 1].Morph(g[c.y][c.x], swaptime)
    }
}


type Block struct {
    color Color;
    state State
    counter uint
    direction Direction
    combo bool;
    morph Block;
}

func (b Block) IsSwappable() {
    return b.morph == nil && b.state != HANG
    return b.state == STILL || b.state == DOWN
}

func (b Block) IsSupport() {
    return b.form != AIR && b.state != DROP
}

func (b *Block) Morph(old Block, duration int) {
    b.counter = duration
    b.morph = old
}

func (b *Block) Tick(under Block) {
    if counter > 0 {
        counter--
        return
    }

    if b.morph != nil {
        b.morph = nil
    }

    switch b.state {
    case STILL:
        if !under.isSupport():
        switch {
        case under.state == HANG:
            b.state = HANG
            b.counter = under.counter
        case under.form == AIR:
            b.state = HANG
            b.counter = hangtime
        }
    case MOVE:
        if under.state == STILL ||
    case HANG:
        b.state
    case DROP:
        switch {
        case under.state == DROP || under.form == AIR:
            b.counter = falltime
        case under.IsEmpty:
            b.counter =

        b.state = DROP
    case LEFT:
    case RIGHT:
    default:
    }
}

type Game struct {
    rows [][]Block
    cursor Cursor
}

func (g *Game) Tick() {
    // decr timers
    // check movement
    // set timers
    // check clear
    // spawn garbage
}

/* NewGame Initializes a game with a viewport of x * y and given amount of
 * block colors */
func NewGame(width, height, colors int) *Game {
    // Set a buffer capacity of twice the height to allow for some garbage
    // and an extra line for spawning blocks
    game := Game{color: colors, rows: make([][]Block, height, 1 + 2 * height)}

    // Initialize the visible rows
    for i, _ := range game.rows {
        game.rows[i] = make([]Block, width)
    }
}

func printf (text string, color Color) {
	fmt.Printf(string(color))
	fmt.Printf(text)
	fmt.Printf(string(BLACK))
}

var FRAME int = 0

func render(game *Game) {
	fmt.Printf("\x1b[H");
	for y, row := range game.grid {
		for i := 0; i < 2; i++ {
			for x, block := range row {
				//fmt.Printf(" ")
				var str string
				if game.cursor.y == y && game.cursor.x == x {
					str = ">   "
				} else if game.cursor.y == y && game.cursor.x == x - 1 {
					str = "   <"
				} else {
					str = "    "
				}
				printf(str, block.color)
			}
			printf("                   ", BLACK)
			fmt.Println("")
		}
		//fmt.Println("")
	}
}


func main() {
	game := Game{}
	game.cursor = Cursor{3, 5}
	game.grid[0][0] = Block{color: RED}
	game.grid[0][1] = Block{color: LIGHT_BLUE}
	game.grid[0][2] = Block{color: GREEN}
	game.grid[0][3] = Block{color: MAGENTA}
	game.grid[0][4] = Block{color: RED}
	game.grid[0][5] = Block{color: YELLOW}
	game.grid[1][3] = Block{color: RED}
	game.grid[1][4] = Block{color: LIGHT_BLUE}
	game.grid[1][5] = Block{color: GREEN}
	game.grid[1][0] = Block{color: MAGENTA}
	game.grid[1][1] = Block{color: RED}
	game.grid[1][2] = Block{color: YELLOW}
	game.grid[2][0] = Block{color: RED}
	game.grid[2][1] = Block{color: LIGHT_BLUE}
	game.grid[2][2] = Block{color: GREEN}
	game.grid[2][3] = Block{color: MAGENTA}
	game.grid[2][4] = Block{color: RED}
	game.grid[2][5] = Block{color: YELLOW}
	game.grid[3][3] = Block{color: RED}
	game.grid[3][4] = Block{color: LIGHT_BLUE}
	game.grid[3][5] = Block{color: GREEN}
	game.grid[3][0] = Block{color: MAGENTA}
	game.grid[3][1] = Block{color: RED}
	game.grid[3][2] = Block{color: YELLOW}
	game.grid[4][0] = Block{color: RED}
	game.grid[4][1] = Block{color: LIGHT_BLUE}
	game.grid[4][2] = Block{color: GREEN}
	game.grid[4][3] = Block{color: MAGENTA}
	game.grid[4][4] = Block{color: RED}
	game.grid[4][5] = Block{color: YELLOW}
	game.grid[5][3] = Block{color: RED}
	game.grid[5][4] = Block{color: LIGHT_BLUE}
	game.grid[5][5] = Block{color: GREEN}
	game.grid[5][0] = Block{color: MAGENTA}
	game.grid[5][1] = Block{color: RED}
	game.grid[5][2] = Block{color: YELLOW}
	game.grid[6][0] = Block{color: RED}
	game.grid[6][1] = Block{color: LIGHT_BLUE}
	game.grid[6][2] = Block{color: GREEN}
	game.grid[6][3] = Block{color: MAGENTA}
	game.grid[6][4] = Block{color: RED}
	game.grid[6][5] = Block{color: YELLOW}
	game.grid[7][3] = Block{color: RED}
	game.grid[7][4] = Block{color: LIGHT_BLUE}
	game.grid[7][5] = Block{color: GREEN}
	game.grid[7][0] = Block{color: MAGENTA}
	game.grid[7][1] = Block{color: RED}
	game.grid[7][2] = Block{color: YELLOW}
	game.grid[8][0] = Block{color: RED}
	game.grid[8][1] = Block{color: LIGHT_BLUE}
	game.grid[8][2] = Block{color: GREEN}
	game.grid[8][3] = Block{color: MAGENTA}
	game.grid[8][4] = Block{color: RED}
	game.grid[8][5] = Block{color: YELLOW}
	game.grid[9][3] = Block{color: RED}
	game.grid[9][4] = Block{color: LIGHT_BLUE}
	game.grid[9][5] = Block{color: GREEN}
	game.grid[9][0] = Block{color: MAGENTA}
	game.grid[9][1] = Block{color: RED}
	game.grid[9][2] = Block{color: YELLOW}
	game.grid[10][0] = Block{color: RED}
	game.grid[10][1] = Block{color: LIGHT_BLUE}
	game.grid[10][2] = Block{color: GREEN}
	game.grid[10][3] = Block{color: MAGENTA}
	game.grid[10][4] = Block{color: RED}
	game.grid[10][5] = Block{color: YELLOW}
	game.grid[11][3] = Block{color: RED}
	game.grid[11][4] = Block{color: LIGHT_BLUE}
	game.grid[11][5] = Block{color: GREEN}
	game.grid[11][0] = Block{color: MAGENTA}
	game.grid[11][1] = Block{color: RED}
	game.grid[11][2] = Block{color: YELLOW}

	for {
		render(&game)
		FRAME += 1
		fmt.Printf("%d", FRAME)
	}

}
