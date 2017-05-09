/* 
 * Character class, encompases and individual character.
 * Holds:
 * - One Character         self.Cargo
 * - Charcter's Index      self.Index
 * - Character's Line      self.Line
 * - Character's column    self.Column
 * - Reference to source   self.Source
 */
class Character {
   public Cargo  char;
   public Index  int;
   public Line   int;
   public Column int;
   public Source &string;

   func Character(c char, i int, l int, col int, source &string){
      this.Cargo = c;
      this.Index = i;
      this.Line = l;
      this.Column = col;
      this.Source = source;
   }
}
