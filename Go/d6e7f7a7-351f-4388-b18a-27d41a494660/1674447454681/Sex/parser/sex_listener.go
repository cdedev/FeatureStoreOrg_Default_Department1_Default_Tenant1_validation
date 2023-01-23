// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447454681/Sex.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SexListener is a complete listener for a parse tree produced by SexParser.
type SexListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSex is called when entering the sex production.
	EnterSex(c *SexContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSex is called when exiting the sex production.
	ExitSex(c *SexContext)
}
