// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675671034100/Alive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Alive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AliveListener is a complete listener for a parse tree produced by AliveParser.
type AliveListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAlive is called when entering the alive production.
	EnterAlive(c *AliveContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAlive is called when exiting the alive production.
	ExitAlive(c *AliveContext)
}
