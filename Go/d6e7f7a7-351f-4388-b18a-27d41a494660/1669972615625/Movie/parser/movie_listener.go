// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972615625/Movie.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Movie

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MovieListener is a complete listener for a parse tree produced by MovieParser.
type MovieListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField0 is called when entering the field0 production.
	EnterField0(c *Field0Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField0 is called when exiting the field0 production.
	ExitField0(c *Field0Context)
}
