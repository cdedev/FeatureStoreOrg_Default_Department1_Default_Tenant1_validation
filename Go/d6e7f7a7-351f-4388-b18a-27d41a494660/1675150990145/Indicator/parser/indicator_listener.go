// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675150990145/Indicator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Indicator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IndicatorListener is a complete listener for a parse tree produced by IndicatorParser.
type IndicatorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIndicator is called when entering the indicator production.
	EnterIndicator(c *IndicatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIndicator is called when exiting the indicator production.
	ExitIndicator(c *IndicatorContext)
}
