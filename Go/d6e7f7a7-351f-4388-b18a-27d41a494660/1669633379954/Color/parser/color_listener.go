// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669633379954/Color.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Color

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ColorListener is a complete listener for a parse tree produced by ColorParser.
type ColorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterColor is called when entering the color production.
	EnterColor(c *ColorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitColor is called when exiting the color production.
	ExitColor(c *ColorContext)
}
