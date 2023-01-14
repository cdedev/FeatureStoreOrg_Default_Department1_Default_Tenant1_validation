// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673665830093/Consumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Consumption

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConsumptionListener is a complete listener for a parse tree produced by ConsumptionParser.
type ConsumptionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConsumption is called when entering the consumption production.
	EnterConsumption(c *ConsumptionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConsumption is called when exiting the consumption production.
	ExitConsumption(c *ConsumptionContext)
}
