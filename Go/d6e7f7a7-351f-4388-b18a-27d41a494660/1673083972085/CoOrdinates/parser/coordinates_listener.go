// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083972085/CoOrdinates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CoOrdinates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CoOrdinatesListener is a complete listener for a parse tree produced by CoOrdinatesParser.
type CoOrdinatesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCoordinates is called when entering the coordinates production.
	EnterCoordinates(c *CoordinatesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCoordinates is called when exiting the coordinates production.
	ExitCoordinates(c *CoordinatesContext)
}
