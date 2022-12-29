// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672287202237/Poutcome.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Poutcome

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PoutcomeListener is a complete listener for a parse tree produced by PoutcomeParser.
type PoutcomeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPoutcome is called when entering the poutcome production.
	EnterPoutcome(c *PoutcomeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPoutcome is called when exiting the poutcome production.
	ExitPoutcome(c *PoutcomeContext)
}
