// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672802581911/WindDirection.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WindDirection

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WindDirectionListener is a complete listener for a parse tree produced by WindDirectionParser.
type WindDirectionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWinddirection is called when entering the winddirection production.
	EnterWinddirection(c *WinddirectionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWinddirection is called when exiting the winddirection production.
	ExitWinddirection(c *WinddirectionContext)
}
