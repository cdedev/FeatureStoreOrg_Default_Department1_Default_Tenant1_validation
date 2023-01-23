// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447401032/Center.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Center

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CenterListener is a complete listener for a parse tree produced by CenterParser.
type CenterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCenter is called when entering the center production.
	EnterCenter(c *CenterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCenter is called when exiting the center production.
	ExitCenter(c *CenterContext)
}
