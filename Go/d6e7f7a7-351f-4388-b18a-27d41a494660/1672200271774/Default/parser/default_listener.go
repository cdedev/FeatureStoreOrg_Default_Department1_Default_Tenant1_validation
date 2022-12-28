// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672200271774/Default.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Default

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DefaultListener is a complete listener for a parse tree produced by DefaultParser.
type DefaultListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)
}
