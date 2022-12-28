// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205224922/CornePoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CornePoint

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CornePointListener is a complete listener for a parse tree produced by CornePointParser.
type CornePointListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCornepoint is called when entering the cornepoint production.
	EnterCornepoint(c *CornepointContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCornepoint is called when exiting the cornepoint production.
	ExitCornepoint(c *CornepointContext)
}
