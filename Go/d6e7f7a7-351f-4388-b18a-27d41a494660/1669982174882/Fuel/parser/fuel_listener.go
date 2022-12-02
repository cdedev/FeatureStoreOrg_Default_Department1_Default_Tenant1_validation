// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982174882/Fuel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fuel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FuelListener is a complete listener for a parse tree produced by FuelParser.
type FuelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFuel is called when entering the fuel production.
	EnterFuel(c *FuelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFuel is called when exiting the fuel production.
	ExitFuel(c *FuelContext)
}
