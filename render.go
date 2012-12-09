package main

const (
	/*
	WHITE Color = "\033[01;37m";
	RED = "\033[22;31m";
	GREEN = "\033[22;32m";
	LIGHT_BLUE = "\033[01;34m";
	MAGENTA = "\033[22;35m";
	YELLOW = "\033[01;33m";
	*/
	WHITE Color = "\033[01;47m"
	RED = "\033[22;41m"
	GREEN = "\033[22;42m"
	LIGHT_BLUE = "\033[01;44m"
	MAGENTA = "\033[22;45m"
	YELLOW = "\033[01;43m"
	BLACK = "\033[22;40m"
)

type Renderer interface {
    Game(g *Game)
    Cursor(x, y int)
    Block(x, y int, b *Block)
}
