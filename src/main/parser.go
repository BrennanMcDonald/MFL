package main

import "fmt"

/* BNF for the language
 * program =               statement {statement} EOF.
 * statement =             printStatement | assignmentStatement
 * printStatement =        "print" expression ";"
 * forStatement =          "while" condition "{" {statement} "}"
 * assignmentStatement =   identifier "=" expression ";"
 * expression =            stringExpression | numberExpression
 * numberExpression =      (numberLiteral | variable) { numberOperator numberExpression }
 * numberOperator =        "+" | "-" | "*" | "/"
 * stringExpression =      (stringLiteral | variable) {"||" stringExpression}
 * condition =             comparable {("and" | "or" | "xor" | "!and" | "!or") comparable}
 * comparitor =            "==" | "<=" | ">=" | "!="
 * comparable =            (numberLiteral | stringLiteral | variable) comparitor (numberLiteral | stringLiteral | variable) { comparitor (numberLiteral | stringLiteral | variable) }
 */

type Parser struct {
	numberOperator []string
	lexer          Lexer
	token          Token
	ast            ASTNode
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
	if p.token.cType == cFoundType {
		p.GetToken()
	} else {
		panic("Was expecting " + cFoundType + " but recieved " + p.token.cType + " at " + p.token.String())
	}
}

func (p *Parser) Parse(sourceText string) ASTNode {
	p.lexer = Lexer{}
	p.numberOperator = []string{"+", "-", "*", "/"}
	p.lexer.Initialize(sourceText)
	p.GetToken()
	p.Program()
	return p.ast
}

func (p *Parser) Program() {
	node := ASTNode{}
	p.Statement(&node)
	for !p.Found(EOF) {
		p.Statement(&node)
	}
	p.Consume(EOF)
	p.ast = node
}

func (p *Parser) Statement(node *ASTNode) {
	/* statement = printStatement | assignmentStatement
	 * printStatement = "print" expression ";"
	 * assignmentStatement = identifier "=" expression ";"
	 */
	if p.Found("print") {
		p.printStatement(node)
	} else if p.Found("while") {
		p.whileStatement(node)
	} else {
		p.assignmentStatement(node)
	}
}

func (p *Parser) Expression(node *ASTNode) {
	if p.Found(NUMBER) {
		var operator ASTNode
		leftSide := ASTNode{p.token, node.level, []ASTNode{}}
		p.GetToken()
		// p.numberLiteral(node)
		for p.FoundOneOf(p.numberOperator) {
			operator = ASTNode{p.token, node.level, []ASTNode{}}
			// node.add(p.token)
			p.GetToken()
			rightSide := ASTNode{p.token, node.level, []ASTNode{}}
			// p.numberExpression(node)
			operator.addNode(leftSide)
			operator.addNode(rightSide)
			p.GetToken()
		}
		node.addNode(operator)
	} else if p.Found(STRING) {
		p.stringLiteral(node)
		for p.Found("||") {
			p.GetToken()
			p.stringExpression(node)
		}
	} else {
		node.add(p.token)
		p.Consume(IDENTIFIER)

		if p.Found("||") {
			for p.Found("||") {
				p.GetToken()
				p.stringExpression(node)
			}
		} else if p.FoundOneOf(p.numberOperator) {
			for p.FoundOneOf(p.numberOperator) {
				node.add(p.token)
				p.GetToken()
				p.numberExpression(node)
			}
		}
	}
}

func (p *Parser) whileStatement(node *ASTNode) {
	headerNode := TokenNode(p.token)
	p.Consume("while")

	expression := TokenNode(Token{CONDITION, CONDITION, p.token.sLine, p.token.sColumn})

	for !(p.Found("{")) {
		headerNode = TokenNode(p.token)
		expression.addNode(headerNode)
		p.GetToken()
	}

	headerNode.addNode(expression)
	body := TokenNode(Token{BODY, BODY, p.token.sLine, p.token.sColumn})

	p.Consume("{")

	for !(p.Found("}")) {
		// TODO: Fix double IDENTIFIER tag
		statementNode := Node()
		p.Statement(&statementNode)
		body.addNode(statementNode)
		fmt.Printf("Body:")
		fmt.Println(p.token)
	}
	headerNode.addNode(body)
	node.addNode(headerNode)

	p.Consume("}")
}

func (p *Parser) printStatement(node *ASTNode) {
	statementNode := TokenNode(p.token)
	p.Consume("print")
	p.Expression(&statementNode)
	node.addNode(statementNode)
	p.Consume(";")
}

func (p *Parser) assignmentStatement(node *ASTNode) {
	idNode := TokenNode(p.token)
	p.Consume(IDENTIFIER)

	opNode := TokenNode(p.token)
	p.Consume("=")

	opNode.addNode(idNode)

	p.Expression(&opNode)
	node.addNode(opNode)

	p.Consume(";")
}

func (p *Parser) stringExpression(node *ASTNode) {
	if p.Found(STRING) {
		node.add(p.token)
		p.GetToken()

		for p.Found("||") {
			p.GetToken()
			p.stringExpression(node)
		}
	} else {
		node.add(p.token)
		p.Consume(IDENTIFIER)
	}

	for p.Found("||") {
		p.GetToken()
		p.stringExpression(node)
	}

}

func (p *Parser) numberExpression(node *ASTNode) {
	if p.Found(NUMBER) {
		p.numberLiteral(node)
	} else {
		node.add(p.token)
		p.Consume(IDENTIFIER)
	}
	for p.FoundOneOf(p.numberOperator) {
		node.add(p.token)
		p.GetToken()
		p.numberExpression(node)
	}
}

func (p *Parser) stringLiteral(node *ASTNode) {
	node.add(p.token)
	p.GetToken()
}

func (p *Parser) numberLiteral(node *ASTNode) {
	node.add(p.token)
	p.GetToken()
}
