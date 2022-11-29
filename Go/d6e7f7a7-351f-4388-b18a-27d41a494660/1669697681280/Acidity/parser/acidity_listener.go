// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697681280/Acidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AcidityListener is a complete listener for a parse tree produced by AcidityParser.
type AcidityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAcidity is called when entering the acidity production.
	EnterAcidity(c *AcidityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAcidity is called when exiting the acidity production.
	ExitAcidity(c *AcidityContext)
}
