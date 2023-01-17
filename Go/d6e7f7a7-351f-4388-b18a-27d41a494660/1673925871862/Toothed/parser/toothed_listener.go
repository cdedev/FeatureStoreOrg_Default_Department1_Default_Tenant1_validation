// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925871862/Toothed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Toothed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ToothedListener is a complete listener for a parse tree produced by ToothedParser.
type ToothedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterToothed is called when entering the toothed production.
	EnterToothed(c *ToothedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitToothed is called when exiting the toothed production.
	ExitToothed(c *ToothedContext)
}
