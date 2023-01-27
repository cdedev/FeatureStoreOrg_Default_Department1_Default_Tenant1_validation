// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674798267059/Turnover.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Turnover

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TurnoverListener is a complete listener for a parse tree produced by TurnoverParser.
type TurnoverListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTurnover is called when entering the turnover production.
	EnterTurnover(c *TurnoverContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTurnover is called when exiting the turnover production.
	ExitTurnover(c *TurnoverContext)
}
