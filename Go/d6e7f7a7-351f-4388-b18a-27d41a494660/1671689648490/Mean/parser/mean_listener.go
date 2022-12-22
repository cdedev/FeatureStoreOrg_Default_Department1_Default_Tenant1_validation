// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689648490/Mean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mean

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MeanListener is a complete listener for a parse tree produced by MeanParser.
type MeanListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMean is called when entering the mean production.
	EnterMean(c *MeanContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMean is called when exiting the mean production.
	ExitMean(c *MeanContext)
}
