// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196033925/RawH2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawH2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RawH2Listener is a complete listener for a parse tree produced by RawH2Parser.
type RawH2Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRawh2 is called when entering the rawh2 production.
	EnterRawh2(c *Rawh2Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRawh2 is called when exiting the rawh2 production.
	ExitRawh2(c *Rawh2Context)
}
