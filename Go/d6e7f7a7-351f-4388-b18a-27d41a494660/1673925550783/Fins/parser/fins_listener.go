// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925550783/Fins.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fins

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FinsListener is a complete listener for a parse tree produced by FinsParser.
type FinsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFins is called when entering the fins production.
	EnterFins(c *FinsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFins is called when exiting the fins production.
	ExitFins(c *FinsContext)
}
