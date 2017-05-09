package main


type Scanner struct{
   source  string
   index   int
   row     int
   column  int
}

func (s *Scanner) Initialize(st string){
   s.index = -1;
   s.row = 0;
   s.column = -1;
   s.source = st;
}

func (s *Scanner) peek() Character {
   if (s.index >= len(s.source)-1){
      char := Character{rune(0), s.index, s.row, s.column, &s.source}
      return char
   }
   c := s.source[s.index+1]
   char := Character{rune(c), s.index, s.row, s.column, &s.source}

   return char
}

func (s *Scanner) getCol() int { return s.column }
func (s *Scanner) getRow() int { return s.row }

func (s *Scanner) get() (bool,Character) {
   s.index += 1;

   if (s.index > 0){
      if (s.source[s.index - 1] == '\n'){
         s.row += 1;
         s.column = -1;
      }
   }

   s.column += 1;
   char := Character{}
   succ := true
   if(s.index >= len(s.source)){
       char = Character{rune(0), s.index, s.row, s.column, &s.source}
       succ = false;
       return succ,char;
   } else {
      c := s.source[s.index]
      char = Character{rune(c), s.index, s.row, s.column, &s.source}
   }
   return succ,char;
}
