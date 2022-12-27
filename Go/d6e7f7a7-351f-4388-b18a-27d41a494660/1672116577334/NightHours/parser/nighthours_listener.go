// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116577334/NightHours.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NightHours

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NightHoursListener is a complete listener for a parse tree produced by NightHoursParser.
type NightHoursListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNighthours is called when entering the nighthours production.
	EnterNighthours(c *NighthoursContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNighthours is called when exiting the nighthours production.
	ExitNighthours(c *NighthoursContext)
}
