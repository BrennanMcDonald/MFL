# MFL
My first attempt at writing a programming language

# Grammar

```
BNF for the language
program =               statement {statement} EOF.
statement =             printStatement | assignmentStatement
printStatement =        "print" expression ";"
forStatement =          "while" condition "{" {statement} "}"
assignmentStatement =   identifier "=" expression ";"
expression =            stringExpression | numberExpression
numberExpression =      (numberLiteral | variable) { numberOperator numberExpression }
numberOperator =        "+" | "-" | "*" | "/"
stringExpression =      (stringLiteral | variable) {"||" stringExpression}
condition =             comparable {("and" | "or" | "xor" | "!and" | "!or") comparable}
comparitor =            "==" | "<=" | ">=" | "!="
comparable =            (numberLiteral | stringLiteral | variable) comparitor (numberLiteral | stringLiteral | variable) { comparitor (numberLiteral | stringLiteral | variable) }
```
## Scanner

`./src/main/scanner.go`

## Lexer

`./src/main/lecer.go`

## Parser

`./src/main/parser.go`
