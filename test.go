package main

import (
	"github.com/jessethegame/colorgrid"
)

func newGame() Game {
	game := NewGame(6, 12, 5)
	game.rows[0][0] = Block{color: colorgrid.RED}
	game.rows[0][1] = Block{color: colorgrid.LIGHT_BLUE}
	game.rows[0][2] = Block{color: colorgrid.GREEN}
	game.rows[0][3] = Block{color: colorgrid.MAGENTA}
	game.rows[0][4] = Block{color: colorgrid.RED}
	game.rows[0][5] = Block{color: colorgrid.YELLOW}
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
	return game
}

func newKeys() keys {
	return keys{
		/*
			up:    termbox.KeyArrowUp,
			down:  termbox.KeyArrowDown,
			left:  termbox.KeyArrowLeft,
			right: termbox.KeyArrowRight,
		*/
		up:    'k',
		down:  'j',
		left:  'h',
		right: 'l',

		top:    'K',
		bottom: 'J',
		home:   'H',
		end:    'L',

		swap: 'f',
		push: 'd',

		quit:  'q',
		pause: 'p',
	}
}

func newPlayer() Player {
	p = Player{
		game: newGame(),
		keys: newKeys(),
	}
	p.cursor = p.game.blocks[p.game.width/2][p.game.height/2]
	return p
}
