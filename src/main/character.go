package main

/*
 * Character class, encompases and individual character.
 * Holds:
 * - One Character         self.Cargo
 * - Charcter's Index      self.Index
 * - Character's Line      self.Line
 * - Character's column    self.Column
 * - Reference to source   self.Source
 */

import (
	"fmt"
)

type Character struct {
	Cargo  rune
	Index  int
	Line   int
	Column int
	Source *string
}

func (c *Character) getCargo() *rune { return &c.Cargo }
func (c *Character) getIndex() *int  { return &c.Index }
func (c *Character) getLine() *int   { return &c.Line }
func (c *Character) getColumn() *int { return &c.Column }

func (c *Character) String() string {
	var cargo string
	if c.Cargo == ' ' {
		cargo = "   Space"
	} else if c.Cargo == '\n' || c.Cargo == '\r' {
		cargo = "   NewLine"
	} else if c.Cargo == '\t' {
		cargo = "   Tab"
	} else if int(c.Cargo) < 1 {
		cargo = "   Endline"
	} else {
		cargo = string(c.Cargo)
	}
	return fmt.Sprintf("%d\t%d %d\t%s", c.Index, c.Line, c.Column, cargo)
}

/*
func main(){
   text := "asdfasd"
   c := Character{'\n', 1, 1, 1, &text};
   fmt.Println(c.String());
}
*/
