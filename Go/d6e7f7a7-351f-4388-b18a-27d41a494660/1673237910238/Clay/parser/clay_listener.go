// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237910238/Clay.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clay

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClayListener is a complete listener for a parse tree produced by ClayParser.
type ClayListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterClay is called when entering the clay production.
	EnterClay(c *ClayContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitClay is called when exiting the clay production.
	ExitClay(c *ClayContext)
}
