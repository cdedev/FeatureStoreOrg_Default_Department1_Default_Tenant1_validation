// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672973411713/Decelerations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Decelerations

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DecelerationsListener is a complete listener for a parse tree produced by DecelerationsParser.
type DecelerationsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDecelerations is called when entering the decelerations production.
	EnterDecelerations(c *DecelerationsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDecelerations is called when exiting the decelerations production.
	ExitDecelerations(c *DecelerationsContext)
}
