// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121353901/Pig.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PigListener is a complete listener for a parse tree produced by PigParser.
type PigListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPig is called when entering the pig production.
	EnterPig(c *PigContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPig is called when exiting the pig production.
	ExitPig(c *PigContext)
}
