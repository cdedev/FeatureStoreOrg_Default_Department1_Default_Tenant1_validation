// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628758415/Temperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Temperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TemperatureListener is a complete listener for a parse tree produced by TemperatureParser.
type TemperatureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTemperature is called when entering the temperature production.
	EnterTemperature(c *TemperatureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTemperature is called when exiting the temperature production.
	ExitTemperature(c *TemperatureContext)
}
