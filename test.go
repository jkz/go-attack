package main

import (
	"github.com/jessethegame/colorgrid"
)

func defaultGame() *Game {
	game := newGame(6, 12, 5)
	game.blocks[0][0] = Block{color: colorgrid.RED}
	game.blocks[1][0] = Block{color: colorgrid.CYAN}
	game.blocks[2][0] = Block{color: colorgrid.GREEN}
	game.blocks[3][0] = Block{color: colorgrid.MAGENTA}
	game.blocks[4][0] = Block{color: colorgrid.RED}
	game.blocks[5][0] = Block{color: colorgrid.YELLOW}
	game.blocks[3][1] = Block{color: colorgrid.RED}
	game.blocks[4][1] = Block{color: colorgrid.CYAN}
	game.blocks[5][1] = Block{color: colorgrid.GREEN}
	game.blocks[0][1] = Block{color: colorgrid.MAGENTA}
	game.blocks[1][1] = Block{color: colorgrid.RED}
	game.blocks[2][1] = Block{color: colorgrid.YELLOW}
	game.blocks[0][2] = Block{color: colorgrid.RED}
	game.blocks[1][2] = Block{color: colorgrid.CYAN}
	game.blocks[2][2] = Block{color: colorgrid.GREEN}
	game.blocks[3][2] = Block{color: colorgrid.MAGENTA}
	game.blocks[4][2] = Block{color: colorgrid.RED}
	game.blocks[5][2] = Block{color: colorgrid.YELLOW}
	game.blocks[3][3] = Block{color: colorgrid.RED}
	game.blocks[4][3] = Block{color: colorgrid.CYAN}
	game.blocks[5][3] = Block{color: colorgrid.GREEN}
	/*
		game.blocks[0][3] = Block{color: colorgrid.MAGENTA}
		game.blocks[1][3] = Block{color: colorgrid.RED}
		game.blocks[2][3] = Block{color: colorgrid.YELLOW}
		game.blocks[0][4] = Block{color: colorgrid.RED}
		game.blocks[1][4] = Block{color: colorgrid.CYAN}
		game.blocks[2][4] = Block{color: colorgrid.GREEN}
		game.blocks[3][4] = Block{color: colorgrid.MAGENTA}
		game.blocks[4][4] = Block{color: colorgrid.RED}
		game.blocks[5][4] = Block{color: colorgrid.YELLOW}
		game.blocks[3][5] = Block{color: colorgrid.RED}
		game.blocks[4][5] = Block{color: colorgrid.CYAN}
		game.blocks[5][5] = Block{color: colorgrid.GREEN}
		game.blocks[0][5] = Block{color: colorgrid.MAGENTA}
		game.blocks[1][5] = Block{color: colorgrid.RED}
		game.blocks[2][5] = Block{color: colorgrid.YELLOW}
		game.blocks[0][6] = Block{color: colorgrid.RED}
		game.blocks[1][6] = Block{color: colorgrid.CYAN}
		game.blocks[2][6] = Block{color: colorgrid.GREEN}
		game.blocks[3][6] = Block{color: colorgrid.MAGENTA}
		game.blocks[4][6] = Block{color: colorgrid.RED}
		game.blocks[5][6] = Block{color: colorgrid.YELLOW}
		game.blocks[3][7] = Block{color: colorgrid.RED}
		game.blocks[4][7] = Block{color: colorgrid.CYAN}
		game.blocks[5][7] = Block{color: colorgrid.GREEN}
		game.blocks[0][7] = Block{color: colorgrid.MAGENTA}
		game.blocks[1][7] = Block{color: colorgrid.RED}
		game.blocks[2][7] = Block{color: colorgrid.YELLOW}
		game.blocks[0][8] = Block{color: colorgrid.RED}
		game.blocks[1][8] = Block{color: colorgrid.CYAN}
		game.blocks[2][8] = Block{color: colorgrid.GREEN}
		game.blocks[3][8] = Block{color: colorgrid.MAGENTA}
		game.blocks[4][8] = Block{color: colorgrid.RED}
		game.blocks[5][8] = Block{color: colorgrid.YELLOW}
		game.blocks[3][9] = Block{color: colorgrid.RED}
		game.blocks[4][9] = Block{color: colorgrid.CYAN}
		game.blocks[5][9] = Block{color: colorgrid.GREEN}
		game.blocks[0][9] = Block{color: colorgrid.MAGENTA}
		game.blocks[1][9] = Block{color: colorgrid.RED}
		game.blocks[2][9] = Block{color: colorgrid.YELLOW}
	*/
	game.blocks[0][10] = Block{color: colorgrid.RED}
	game.blocks[1][10] = Block{color: colorgrid.CYAN}
	game.blocks[2][10] = Block{color: colorgrid.GREEN}
	game.blocks[3][10] = Block{color: colorgrid.MAGENTA}
	game.blocks[4][10] = Block{color: colorgrid.RED}
	game.blocks[5][10] = Block{color: colorgrid.YELLOW}
	game.blocks[3][11] = Block{color: colorgrid.RED}
	game.blocks[4][11] = Block{color: colorgrid.CYAN}
	game.blocks[5][11] = Block{color: colorgrid.GREEN}
	game.blocks[0][11] = Block{color: colorgrid.MAGENTA}
	game.blocks[1][11] = Block{color: colorgrid.RED}
	game.blocks[2][11] = Block{color: colorgrid.YELLOW}
	return game
}

func defaultKeys() keys {
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

		swap: 32,
		push: 'd',

		quit:  'q',
		pause: 'p',
	}
}

func defaultPlayer() Player {
	p := Player{
		game: defaultGame(),
		keys: defaultKeys(),
	}
	p.cursor = &(p.game.blocks[p.game.width/2][p.game.height/2])
	return p
}
