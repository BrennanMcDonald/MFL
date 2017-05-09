package main

type Lexer struct {
   c1       string
   c2       string
   scan     Scanner
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (l *Lexer) Initialize(s string) {
   l.scan.Initialize(s)
   l.GetChar()
}

func (l *Lexer) GetNextToken() Token {
   for (contains(WHITESPACE_CHARS, l.c1) || l.c2 == "/*") {
      /*   Handle whitespace
       *   and comments in the language
       *   return nothing
       */
      if contains(WHITESPACE_CHARS, l.c1){
         token := Token{l.c1, WHITESPACE, l.scan.getCol(), l.scan.getRow()}
         for contains(WHITESPACE_CHARS, l.c1) {
            l.GetChar()
            token.Cargo += l.c1
         }
      }

      if l.c2 == "/*" {
         // Handles comments, slash astrisk style
         token := Token{l.c1, COMMENT, l.scan.getCol(), l.scan.getRow()}
         l.GetChar()
         l.GetChar()
         for l.c2 != "*/" {
            // May fuck up here
            if(l.scan.peek().Cargo == rune(0)){
               panic("End of text found before end of comment")
            }
            l.GetChar()
            token.Cargo += l.c1
         }
         l.GetChar()
         l.GetChar()
      }

   } // End Whitespace Handling

   token := Token{}
   token.sLine = l.scan.getRow()
   token.sColumn = l.scan.getCol()
   if (l.c1 == string(rune(0))) {
      token.cType = EOF;
      return token;
   }

   if (contains(IDENTIFIER_STARTCHARS, l.c1)){
      token.cType = IDENTIFIER;
      token.Cargo = l.c1
      l.GetChar()
      for contains(IDENTIFIER_CHARS, l.c1){
         token.Cargo += l.c1
         l.GetChar()
      }
      if(contains(Keywords,token.Cargo)){
         token.cType = token.Cargo
      }
      return token
   }

   if (contains(NUMBER_STARTCHARS, l.c1)){
      token.cType = NUMBER;
      token.Cargo = l.c1
      l.GetChar()
      for contains(NUMBER_CHARS, l.c1){
         token.Cargo += l.c1
         l.GetChar()
      }
      return token
   }

   if (contains(STRING_STARTCHARS, l.c1)){
      start_char := l.c1
      token.cType = STRING
      token.Cargo = l.c1
      l.GetChar()
      for l.c1 != start_char {
         if(l.scan.peek().Cargo == rune(0)){
            panic("End of text found before end of comment")
         }
         token.Cargo += l.c1
         l.GetChar()
      }
      token.Cargo += l.c1
      l.GetChar()
      return token
   }

   if (contains(TwoCharacterSymbols, l.c2)){
      token.Cargo = l.c2
      token.cType = "Symbol"
      l.GetChar()
      l.GetChar()
      return token;
   }

   if (contains(OneCharacterSymbols, l.c1)){
      token.Cargo = l.c1
      token.cType = "Symbol"
      l.GetChar()
      return token;
   }
   if (l.c1 == string(rune(0))){
      token.Cargo = EOF;
   }
   return token;

}

func (l *Lexer) GetChar() {
   _,ch := l.scan.get()
   l.c1 = string(ch.Cargo)
   l.c2 = l.c1 + string(l.scan.peek().Cargo)
}
