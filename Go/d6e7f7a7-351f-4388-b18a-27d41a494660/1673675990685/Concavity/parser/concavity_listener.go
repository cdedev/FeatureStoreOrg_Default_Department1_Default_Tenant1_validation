// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673675990685/Concavity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concavity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConcavityListener is a complete listener for a parse tree produced by ConcavityParser.
type ConcavityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConcavity is called when entering the concavity production.
	EnterConcavity(c *ConcavityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConcavity is called when exiting the concavity production.
	ExitConcavity(c *ConcavityContext)
}