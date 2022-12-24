// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867422033/ExchUsd.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ExchUsd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExchUsdListener is a complete listener for a parse tree produced by ExchUsdParser.
type ExchUsdListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExchusd is called when entering the exchusd production.
	EnterExchusd(c *ExchusdContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExchusd is called when exiting the exchusd production.
	ExitExchusd(c *ExchusdContext)
}
