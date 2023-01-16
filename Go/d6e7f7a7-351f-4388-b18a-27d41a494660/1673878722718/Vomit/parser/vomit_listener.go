// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878722718/Vomit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vomit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VomitListener is a complete listener for a parse tree produced by VomitParser.
type VomitListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVomit is called when entering the vomit production.
	EnterVomit(c *VomitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVomit is called when exiting the vomit production.
	ExitVomit(c *VomitContext)
}
