// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656626453/Airquality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airquality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AirqualityListener is a complete listener for a parse tree produced by AirqualityParser.
type AirqualityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAirquality is called when entering the airquality production.
	EnterAirquality(c *AirqualityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAirquality is called when exiting the airquality production.
	ExitAirquality(c *AirqualityContext)
}
