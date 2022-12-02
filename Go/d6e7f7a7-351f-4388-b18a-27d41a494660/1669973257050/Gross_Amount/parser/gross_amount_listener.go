// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669973257050/Gross_Amount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gross_Amount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Gross_AmountListener is a complete listener for a parse tree produced by Gross_AmountParser.
type Gross_AmountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGross_amount is called when entering the gross_amount production.
	EnterGross_amount(c *Gross_amountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGross_amount is called when exiting the gross_amount production.
	ExitGross_amount(c *Gross_amountContext)
}
