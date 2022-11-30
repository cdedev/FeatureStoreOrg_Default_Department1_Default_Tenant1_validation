// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669787921220/Barometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BarometerListener is a complete listener for a parse tree produced by BarometerParser.
type BarometerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBarometer is called when entering the barometer production.
	EnterBarometer(c *BarometerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBarometer is called when exiting the barometer production.
	ExitBarometer(c *BarometerContext)
}
