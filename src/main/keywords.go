package main

import (
	"strings"
)

var lower_letters = "a b c d e f g h i j k l m n o p q r s t u v w x y z"
var letters = lower_letters + "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z"
var digits = "1 2 3 4 5 6 7 8 9 0"

var Keywords = strings.Split("for if then else elif endif while loop endloop print return exit", " ")

var OneCharacterSymbols = strings.Split("= ( ) < > / * + - ! & .  ; { }", " ")

var TwoCharacterSymbols = strings.Split("== <= >= <> != ++ ** -- += -= ||", " ")

var IDENTIFIER_STARTCHARS = strings.Split(letters, " ")
var IDENTIFIER_CHARS = strings.Split(letters+digits+" _", " ")

var NUMBER_STARTCHARS = strings.Split(digits, " ")
var NUMBER_CHARS = strings.Split(digits+" .", " ")

var STRING_STARTCHARS = []string{"\"", "'"}
var WHITESPACE_CHARS = []string{" ", "\t", "\n", "\r"}

var STRING = "String"
var IDENTIFIER = "Identifier"
var NUMBER = "Number"
var WHITESPACE = "Whitespace"
var COMMENT = "Comment"
var CONDITION = "Condition"
var BODY = "Body"
var EOF = "Eof"
