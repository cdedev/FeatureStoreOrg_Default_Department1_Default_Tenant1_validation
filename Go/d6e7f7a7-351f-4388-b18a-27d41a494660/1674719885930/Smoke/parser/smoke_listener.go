// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719885930/Smoke.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Smoke

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SmokeListener is a complete listener for a parse tree produced by SmokeParser.
type SmokeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSmoke is called when entering the smoke production.
	EnterSmoke(c *SmokeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSmoke is called when exiting the smoke production.
	ExitSmoke(c *SmokeContext)
}
