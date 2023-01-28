// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881836211/Frozen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Frozen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FrozenListener is a complete listener for a parse tree produced by FrozenParser.
type FrozenListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFrozen is called when entering the frozen production.
	EnterFrozen(c *FrozenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFrozen is called when exiting the frozen production.
	ExitFrozen(c *FrozenContext)
}
