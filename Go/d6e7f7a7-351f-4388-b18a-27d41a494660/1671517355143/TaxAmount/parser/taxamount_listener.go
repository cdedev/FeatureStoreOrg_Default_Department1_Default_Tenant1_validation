// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517355143/TaxAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TaxAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TaxAmountListener is a complete listener for a parse tree produced by TaxAmountParser.
type TaxAmountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTaxamount is called when entering the taxamount production.
	EnterTaxamount(c *TaxamountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTaxamount is called when exiting the taxamount production.
	ExitTaxamount(c *TaxamountContext)
}
