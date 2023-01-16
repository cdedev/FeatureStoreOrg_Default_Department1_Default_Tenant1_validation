// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868556010/Carat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CaratListener is a complete listener for a parse tree produced by CaratParser.
type CaratListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCarat is called when entering the carat production.
	EnterCarat(c *CaratContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCarat is called when exiting the carat production.
	ExitCarat(c *CaratContext)
}
