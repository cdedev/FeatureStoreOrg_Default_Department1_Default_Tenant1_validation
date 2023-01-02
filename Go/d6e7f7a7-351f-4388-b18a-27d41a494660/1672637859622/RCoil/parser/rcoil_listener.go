// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637859622/RCoil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RCoil

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RCoilListener is a complete listener for a parse tree produced by RCoilParser.
type RCoilListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRcoil is called when entering the rcoil production.
	EnterRcoil(c *RcoilContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRcoil is called when exiting the rcoil production.
	ExitRcoil(c *RcoilContext)
}
