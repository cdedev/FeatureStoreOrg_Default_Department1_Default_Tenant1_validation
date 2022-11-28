// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669637492645/location.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // location

import "github.com/antlr/antlr4/runtime/Go/antlr"

// locationListener is a complete listener for a parse tree produced by locationParser.
type locationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLocation is called when entering the location production.
	EnterLocation(c *LocationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLocation is called when exiting the location production.
	ExitLocation(c *LocationContext)
}
