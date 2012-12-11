package main

type Controls struct {
    Up, Down, Left, Right, Swap, Boost, Pause int
}

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

