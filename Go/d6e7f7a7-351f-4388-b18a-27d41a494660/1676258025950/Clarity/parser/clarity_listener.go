// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676258025950/Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClarityListener is a complete listener for a parse tree produced by ClarityParser.
type ClarityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterClarity is called when entering the clarity production.
	EnterClarity(c *ClarityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitClarity is called when exiting the clarity production.
	ExitClarity(c *ClarityContext)
}
