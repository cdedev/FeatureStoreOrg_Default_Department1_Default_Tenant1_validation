// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693130089/FuelType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FuelTypeListener is a complete listener for a parse tree produced by FuelTypeParser.
type FuelTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFueltype is called when entering the fueltype production.
	EnterFueltype(c *FueltypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFueltype is called when exiting the fueltype production.
	ExitFueltype(c *FueltypeContext)
}
