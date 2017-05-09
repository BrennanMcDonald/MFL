package main

// A Token object is the kind of thing that the Lexer returns.
// It holds:
// - the text of the token (self.cargo)
// - the type of token that it is
// - the line number and column index where the token starts

import (
   "fmt"
)

type Token struct {
   Cargo    string
   cType    string
   sLine    int
   sColumn  int
}

func (t Token) String() string{
   return fmt.Sprintf("%d %d\t%s\t%s", t.sLine, t.sColumn, t.Cargo, t.cType)
}
