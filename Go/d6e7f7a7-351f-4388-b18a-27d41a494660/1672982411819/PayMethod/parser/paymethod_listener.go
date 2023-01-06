// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672982411819/PayMethod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PayMethod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PayMethodListener is a complete listener for a parse tree produced by PayMethodParser.
type PayMethodListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPaymethod is called when entering the paymethod production.
	EnterPaymethod(c *PaymethodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPaymethod is called when exiting the paymethod production.
	ExitPaymethod(c *PaymethodContext)
}
