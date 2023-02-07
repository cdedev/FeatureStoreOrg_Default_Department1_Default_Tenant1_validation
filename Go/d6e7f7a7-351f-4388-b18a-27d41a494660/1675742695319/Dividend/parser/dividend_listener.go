// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742695319/Dividend.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dividend

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DividendListener is a complete listener for a parse tree produced by DividendParser.
type DividendListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDividend is called when entering the dividend production.
	EnterDividend(c *DividendContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDividend is called when exiting the dividend production.
	ExitDividend(c *DividendContext)
}
