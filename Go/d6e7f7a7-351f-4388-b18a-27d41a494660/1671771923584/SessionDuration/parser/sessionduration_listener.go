// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771923584/SessionDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SessionDuration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SessionDurationListener is a complete listener for a parse tree produced by SessionDurationParser.
type SessionDurationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSessionduration is called when entering the sessionduration production.
	EnterSessionduration(c *SessiondurationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSessionduration is called when exiting the sessionduration production.
	ExitSessionduration(c *SessiondurationContext)
}
