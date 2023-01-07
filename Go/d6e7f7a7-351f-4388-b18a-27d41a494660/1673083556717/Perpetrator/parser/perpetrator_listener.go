// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083556717/Perpetrator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Perpetrator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PerpetratorListener is a complete listener for a parse tree produced by PerpetratorParser.
type PerpetratorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPerpetrator is called when entering the perpetrator production.
	EnterPerpetrator(c *PerpetratorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPerpetrator is called when exiting the perpetrator production.
	ExitPerpetrator(c *PerpetratorContext)
}
