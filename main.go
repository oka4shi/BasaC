package main

import ()

type TokenType int

type Token struct {
	TokenType TokenType
	Value     string
	Raw       string
}

type Calc struct {
	tokens    []*Token
	variables map[string]int
	stack     []string
	generator *TokenGenerator
}

type TokenGenerator struct {
	source       []rune
	maxIntex     int
	currentIndex int
}

func main() {

}
