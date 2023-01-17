// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925818010/Tail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tail

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TailListener is a complete listener for a parse tree produced by TailParser.
type TailListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTail is called when entering the tail production.
	EnterTail(c *TailContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTail is called when exiting the tail production.
	ExitTail(c *TailContext)
}
