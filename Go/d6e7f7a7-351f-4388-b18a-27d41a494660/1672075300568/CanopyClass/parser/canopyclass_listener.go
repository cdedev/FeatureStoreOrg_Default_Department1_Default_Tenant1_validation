// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672075300568/CanopyClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CanopyClass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CanopyClassListener is a complete listener for a parse tree produced by CanopyClassParser.
type CanopyClassListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCanopyclass is called when entering the canopyclass production.
	EnterCanopyclass(c *CanopyclassContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCanopyclass is called when exiting the canopyclass production.
	ExitCanopyclass(c *CanopyclassContext)
}
