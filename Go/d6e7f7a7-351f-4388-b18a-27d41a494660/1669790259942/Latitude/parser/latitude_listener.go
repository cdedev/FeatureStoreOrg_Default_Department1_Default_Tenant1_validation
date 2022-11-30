// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669790259942/Latitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Latitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LatitudeListener is a complete listener for a parse tree produced by LatitudeParser.
type LatitudeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLatitude is called when entering the latitude production.
	EnterLatitude(c *LatitudeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLatitude is called when exiting the latitude production.
	ExitLatitude(c *LatitudeContext)
}
