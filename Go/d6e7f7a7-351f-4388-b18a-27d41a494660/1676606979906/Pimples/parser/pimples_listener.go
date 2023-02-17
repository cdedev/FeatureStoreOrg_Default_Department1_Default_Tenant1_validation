// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606979906/Pimples.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pimples

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PimplesListener is a complete listener for a parse tree produced by PimplesParser.
type PimplesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPimples is called when entering the pimples production.
	EnterPimples(c *PimplesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPimples is called when exiting the pimples production.
	ExitPimples(c *PimplesContext)
}
