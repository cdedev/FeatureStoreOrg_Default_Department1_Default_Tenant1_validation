// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601979388/StopDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StopDuration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StopDurationListener is a complete listener for a parse tree produced by StopDurationParser.
type StopDurationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStopduration is called when entering the stopduration production.
	EnterStopduration(c *StopdurationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStopduration is called when exiting the stopduration production.
	ExitStopduration(c *StopdurationContext)
}
