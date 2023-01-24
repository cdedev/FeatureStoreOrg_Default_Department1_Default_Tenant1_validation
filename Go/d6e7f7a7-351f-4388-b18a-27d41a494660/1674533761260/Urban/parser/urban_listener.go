// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674533761260/Urban.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Urban

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UrbanListener is a complete listener for a parse tree produced by UrbanParser.
type UrbanListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUrban is called when entering the urban production.
	EnterUrban(c *UrbanContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUrban is called when exiting the urban production.
	ExitUrban(c *UrbanContext)
}
