// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675222431329/TaxType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TaxType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TaxTypeListener is a complete listener for a parse tree produced by TaxTypeParser.
type TaxTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTaxtype is called when entering the taxtype production.
	EnterTaxtype(c *TaxtypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTaxtype is called when exiting the taxtype production.
	ExitTaxtype(c *TaxtypeContext)
}
