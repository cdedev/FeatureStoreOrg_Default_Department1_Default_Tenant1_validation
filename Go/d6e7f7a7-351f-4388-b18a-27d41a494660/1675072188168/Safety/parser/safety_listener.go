// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072188168/Safety.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Safety

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SafetyListener is a complete listener for a parse tree produced by SafetyParser.
type SafetyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSafety is called when entering the safety production.
	EnterSafety(c *SafetyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSafety is called when exiting the safety production.
	ExitSafety(c *SafetyContext)
}
