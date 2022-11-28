// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623262522/Longitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Longitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LongitudeListener is a complete listener for a parse tree produced by LongitudeParser.
type LongitudeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLongitude is called when entering the longitude production.
	EnterLongitude(c *LongitudeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLongitude is called when exiting the longitude production.
	ExitLongitude(c *LongitudeContext)
}
