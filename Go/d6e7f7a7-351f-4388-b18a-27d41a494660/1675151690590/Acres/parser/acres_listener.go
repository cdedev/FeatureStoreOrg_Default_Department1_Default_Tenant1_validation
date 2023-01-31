// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675151690590/Acres.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acres

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AcresListener is a complete listener for a parse tree produced by AcresParser.
type AcresListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAcres is called when entering the acres production.
	EnterAcres(c *AcresContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAcres is called when exiting the acres production.
	ExitAcres(c *AcresContext)
}
