// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669625040586/Tax_Amount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tax_Amount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Tax_AmountListener is a complete listener for a parse tree produced by Tax_AmountParser.
type Tax_AmountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTax_amount is called when entering the tax_amount production.
	EnterTax_amount(c *Tax_amountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTax_amount is called when exiting the tax_amount production.
	ExitTax_amount(c *Tax_amountContext)
}
