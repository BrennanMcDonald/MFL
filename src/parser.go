package main

/* BNF for the language
 * program =               statement {statement} EOF.
 * statement =             printStatement | assignmentStatement
 * printStatement =        "print" expression ";"
 * assignmentStatement =   identifier "=" expression ";"
 * expression =            stringExpression | numberExpression
 * numberExpression =      (numberLiteral | variable) { numberOperator numberExpression }
 * numberOperator =        "+" | "-" | "*" | "/"
 * stringExpression =      (stringLiteral | variable) {"||" stringExpression}
*/

type Parser struct {
   numberOperator    []string
   lexer             Lexer
   token             Token
   ast               ASTNode
}

func (p *Parser) GetToken() {
   p.token = p.lexer.GetNextToken()
}

func (p *Parser) Found(cFoundType string) bool {
   return p.token.cType == cFoundType
}

func (p *Parser) FoundOneOf(cFoundTypes []string) bool {
   return contains(cFoundTypes, p.token.cType)
}

func (p *Parser) Consume(cFoundType string) {
   if (p.token.cType == cFoundType){
      p.GetToken()
   } else {
      panic ("Was expecting " + p.token.cType + " but recieved " + cFoundType + " at " + p.token.String())
   }
}

func (p *Parser) Parse(sourceText string) ASTNode{
   p.lexer = Lexer{}
   p.numberOperator = []string{"+","-","*","/"}
   p.lexer.Initialize(sourceText)
   p.GetToken()
   p.Program()
   return p.ast
}

func (p *Parser) Program() {
   node := ASTNode{}
   p.Statement(node)
   for !p.Found(EOF) {
      p.Statement(node)
   }
   p.Consume(EOF)
   p.ast = node
}

func (p *Parser) Statement(node ASTNode) {
   /* statement = printStatement | assignmentStatement
    * printStatement = "print" expression ";"
    * assignmentStatement = identifier "=" expression ";"
    */
   if (p.Found("print")){
      p.printStatement(&node)
   } else {
      p.assignmentStatement(&node)
   }
}

func (p *Parser) Expression(node *ASTNode) {
   if (p.Found(NUMBER)) {
      p.numberLiteral(node)
      for p.FoundOneOf(p.numberOperator){
         node.add(p.token)
         p.GetToken()
         p.numberExpression(node)
      }
   } else if (p.Found(STRING)){
      p.stringLiteral(node)
      for p.Found("||") {
         p.GetToken()
         p.stringExpression(node)
      }
   } else {
      node.add(p.token)
      p.Consume(IDENTIFIER)

      if (p.Found("||")){
         for (p.Found("||")){
            p.GetToken()
            p.stringExpression(node)
         }
      } else if (p.FoundOneOf(p.numberOperator)) {
         for (p.FoundOneOf(p.numberOperator)){
            node.add(p.token)
            p.GetToken()
            p.numberExpression(node)
         }
      }
   }
}

func (p *Parser) printStatement(node *ASTNode) {
   statementNode := TokenNode(p.token)
   p.Consume("print")
   node.addNode(statementNode)
   p.Expression(&statementNode)
   p.Consume(";")
}

func (p *Parser) assignmentStatement(node *ASTNode) {
   idNode := TokenNode(p.token)
   p.Consume(IDENTIFIER)

   opNode := TokenNode(p.token)
   p.Consume("=")

   opNode.addNode(idNode)
   node.addNode(opNode)

   p.Expression(&opNode)
   p.Consume(";")
}

func (p *Parser) stringExpression(node *ASTNode){
   if (p.Found(STRING)){
      node.add(p.token)
      p.GetToken()

      for (p.Found("||")){
         p.GetToken()
         p.stringExpression(node)
      }
   } else {
      node.add(p.token)
      p.Consume(IDENTIFIER)
   }

   for(p.Found("||")){
      p.GetToken()
      p.stringExpression(node)
   }

}

func (p *Parser) numberExpression(node *ASTNode){
   if (p.Found(NUMBER)){
      p.numberLiteral(node)
   } else {
      node.add(p.token)
      p.Consume(IDENTIFIER)
   }
   for p.FoundOneOf(p.numberOperator){
      node.add(p.token)
      p.GetToken()
      p.numberExpression(node)
   }
}

func (p *Parser) stringLiteral(node *ASTNode){
   node.add(p.token)
   p.GetToken()
}

func (p *Parser) numberLiteral(node *ASTNode){
   node.add(p.token)
   p.GetToken()
}
