// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868747376/Cut.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cut

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CutListener is a complete listener for a parse tree produced by CutParser.
type CutListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCut is called when entering the cut production.
	EnterCut(c *CutContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCut is called when exiting the cut production.
	ExitCut(c *CutContext)
}
