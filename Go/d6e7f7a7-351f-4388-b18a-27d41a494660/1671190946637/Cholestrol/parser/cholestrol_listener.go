// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671190946637/Cholestrol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholestrol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CholestrolListener is a complete listener for a parse tree produced by CholestrolParser.
type CholestrolListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCholestrol is called when entering the cholestrol production.
	EnterCholestrol(c *CholestrolContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCholestrol is called when exiting the cholestrol production.
	ExitCholestrol(c *CholestrolContext)
}
