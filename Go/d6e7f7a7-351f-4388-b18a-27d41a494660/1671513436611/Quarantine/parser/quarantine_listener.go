// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671513436611/Quarantine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quarantine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QuarantineListener is a complete listener for a parse tree produced by QuarantineParser.
type QuarantineListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterQuarantine is called when entering the quarantine production.
	EnterQuarantine(c *QuarantineContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitQuarantine is called when exiting the quarantine production.
	ExitQuarantine(c *QuarantineContext)
}
