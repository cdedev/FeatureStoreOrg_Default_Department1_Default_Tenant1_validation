// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789488857/Magnitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Magnitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MagnitudeListener is a complete listener for a parse tree produced by MagnitudeParser.
type MagnitudeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMagnitude is called when entering the magnitude production.
	EnterMagnitude(c *MagnitudeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMagnitude is called when exiting the magnitude production.
	ExitMagnitude(c *MagnitudeContext)
}
