// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193430361/Genre.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Genre

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GenreListener is a complete listener for a parse tree produced by GenreParser.
type GenreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGenre is called when entering the genre production.
	EnterGenre(c *GenreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGenre is called when exiting the genre production.
	ExitGenre(c *GenreContext)
}
