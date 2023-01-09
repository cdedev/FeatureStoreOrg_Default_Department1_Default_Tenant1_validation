// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238889397/Sensitive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sensitive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SensitiveListener is a complete listener for a parse tree produced by SensitiveParser.
type SensitiveListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSensitive is called when entering the sensitive production.
	EnterSensitive(c *SensitiveContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSensitive is called when exiting the sensitive production.
	ExitSensitive(c *SensitiveContext)
}
