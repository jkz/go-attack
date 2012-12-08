package main

import "fmt"

type State int
type Form int

const (
	AIR Form = iota;
	SOLID;
	GARBAGE;
)

const (
	STATIC State = iota;
	HANG;
	FALL;
	SWAP;
	CLEAR;

	TO_AIR;
	TO_SOLID;
	TO_GARBAGE;
)

type Cursor struct {
	x, y int;
}

type Direction int
var LEFT, STILL, RIGHT = -1, 0, 1



type Color string;

const (
	/*
	WHITE Color = "\033[01;37m";
	RED = "\033[22;31m";
	GREEN = "\033[22;32m";
	LIGHT_BLUE = "\033[01;34m";
	MAGENTA = "\033[22;35m";
	YELLOW = "\033[01;33m";
	*/
	WHITE Color = "\033[01;47m";
	RED = "\033[22;41m";
	GREEN = "\033[22;42m";
	LIGHT_BLUE = "\033[01;44m";
	MAGENTA = "\033[22;45m";
	YELLOW = "\033[01;43m";
	BLACK = "\033[22;40m"
)

type Block struct {
	state State;
	counter uint;
	direction Direction;
	color Color;
}

type Game struct {
	grid [12][6]Block;
	cursor Cursor;
}

func printf (text string, color Color) {
	fmt.Printf(string(color))
	fmt.Printf(text)
	fmt.Printf(string(BLACK))
}

var FRAME int = 0;

func render(game *Game) {
	fmt.Printf("\x1b[H");
	for y, row := range game.grid {
		for i := 0; i < 2; i++ {
			for x, block := range row {
				//fmt.Printf(" ")
				var str string;
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
		render(&game);
		FRAME += 1
		fmt.Printf("%d", FRAME)
	}

}
