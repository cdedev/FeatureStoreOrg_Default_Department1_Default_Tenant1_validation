// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118129023/ApparentTemperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ApparentTemperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ApparentTemperatureListener is a complete listener for a parse tree produced by ApparentTemperatureParser.
type ApparentTemperatureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterApparenttemperature is called when entering the apparenttemperature production.
	EnterApparenttemperature(c *ApparenttemperatureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitApparenttemperature is called when exiting the apparenttemperature production.
	ExitApparenttemperature(c *ApparenttemperatureContext)
}
