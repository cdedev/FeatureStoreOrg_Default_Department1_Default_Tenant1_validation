// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376933873/ProdYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ProdYear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProdYearListener is a complete listener for a parse tree produced by ProdYearParser.
type ProdYearListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProdyear is called when entering the prodyear production.
	EnterProdyear(c *ProdyearContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProdyear is called when exiting the prodyear production.
	ExitProdyear(c *ProdyearContext)
}
