// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675240859030/Penality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Penality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PenalityListener is a complete listener for a parse tree produced by PenalityParser.
type PenalityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPenality is called when entering the penality production.
	EnterPenality(c *PenalityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPenality is called when exiting the penality production.
	ExitPenality(c *PenalityContext)
}
