// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769755680/GrossAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GrossAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrossAmountListener is a complete listener for a parse tree produced by GrossAmountParser.
type GrossAmountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGrossamount is called when entering the grossamount production.
	EnterGrossamount(c *GrossamountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGrossamount is called when exiting the grossamount production.
	ExitGrossamount(c *GrossamountContext)
}
