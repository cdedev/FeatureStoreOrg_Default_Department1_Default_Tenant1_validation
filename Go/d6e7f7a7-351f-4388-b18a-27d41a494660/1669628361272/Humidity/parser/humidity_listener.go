// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628361272/Humidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Humidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HumidityListener is a complete listener for a parse tree produced by HumidityParser.
type HumidityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHumidity is called when entering the humidity production.
	EnterHumidity(c *HumidityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHumidity is called when exiting the humidity production.
	ExitHumidity(c *HumidityContext)
}
