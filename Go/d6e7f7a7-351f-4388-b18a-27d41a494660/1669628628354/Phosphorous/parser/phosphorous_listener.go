// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628628354/Phosphorous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phosphorous

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhosphorousListener is a complete listener for a parse tree produced by PhosphorousParser.
type PhosphorousListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPhosphorous is called when entering the phosphorous production.
	EnterPhosphorous(c *PhosphorousContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPhosphorous is called when exiting the phosphorous production.
	ExitPhosphorous(c *PhosphorousContext)
}
