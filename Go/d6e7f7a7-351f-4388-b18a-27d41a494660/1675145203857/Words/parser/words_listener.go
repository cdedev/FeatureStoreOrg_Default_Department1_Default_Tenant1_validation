// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145203857/Words.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Words

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WordsListener is a complete listener for a parse tree produced by WordsParser.
type WordsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWords is called when entering the words production.
	EnterWords(c *WordsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWords is called when exiting the words production.
	ExitWords(c *WordsContext)
}
