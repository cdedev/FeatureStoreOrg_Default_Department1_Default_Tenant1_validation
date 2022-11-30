// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669788081665/Direction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Direction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DirectionListener is a complete listener for a parse tree produced by DirectionParser.
type DirectionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)
}
