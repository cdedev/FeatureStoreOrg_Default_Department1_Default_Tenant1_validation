// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671768897139/FuelConsumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelConsumption

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FuelConsumptionListener is a complete listener for a parse tree produced by FuelConsumptionParser.
type FuelConsumptionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFuelconsumption is called when entering the fuelconsumption production.
	EnterFuelconsumption(c *FuelconsumptionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFuelconsumption is called when exiting the fuelconsumption production.
	ExitFuelconsumption(c *FuelconsumptionContext)
}
