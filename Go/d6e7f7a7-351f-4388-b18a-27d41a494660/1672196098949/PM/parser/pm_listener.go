// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196098949/PM.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PM

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PMListener is a complete listener for a parse tree produced by PMParser.
type PMListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPm is called when entering the pm production.
	EnterPm(c *PmContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPm is called when exiting the pm production.
	ExitPm(c *PmContext)
}
