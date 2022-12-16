// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206363560/Cured.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cured

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CuredListener is a complete listener for a parse tree produced by CuredParser.
type CuredListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCured is called when entering the cured production.
	EnterCured(c *CuredContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCured is called when exiting the cured production.
	ExitCured(c *CuredContext)
}
