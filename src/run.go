package main


import (
   "io/ioutil"
   "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){

   dat, err := ioutil.ReadFile("code.mfl")
   check(err)

   var l Lexer
   var p Parser
   l.Initialize(string(dat))
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
