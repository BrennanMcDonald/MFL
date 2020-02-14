package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("code.mfl")
	check(err)

	// var l Lexer
	// l.Initialize(string(dat))
	//
	// for {
	//    h := l.GetNextToken()
	//    fmt.Print(h)
	//    if (h.cType == EOF){
	//       break;
	//    }
	// }

	var p Parser
	ast := p.Parse(string(dat))
	fmt.Println(ast.String())

	// var s Scanner;
	// s.Initialize(string(dat))
	// arr := make([]Character, len(string(dat)))
	// for {
	//    _,char := s.get()
	//    fmt.Println(char.String())
	//    arr = append(arr,char)
	//    if (char.Cargo == rune(0)){
	//       break;
	//    }
	// }
}
