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
   l.Initialize(string(dat))

   for {
      h := l.GetNextToken()
      fmt.Println(h)
      if (h.cType == EOF){
         break;
      }
   }

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
