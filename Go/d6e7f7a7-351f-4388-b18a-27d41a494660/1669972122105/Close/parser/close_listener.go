// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972122105/Close.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Close

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CloseListener is a complete listener for a parse tree produced by CloseParser.
type CloseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterClose is called when entering the close production.
	EnterClose(c *CloseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitClose is called when exiting the close production.
	ExitClose(c *CloseContext)
}
