// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740638061/DollarValue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DollarValue

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DollarValueListener is a complete listener for a parse tree produced by DollarValueParser.
type DollarValueListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDollarvalue is called when entering the dollarvalue production.
	EnterDollarvalue(c *DollarvalueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDollarvalue is called when exiting the dollarvalue production.
	ExitDollarvalue(c *DollarvalueContext)
}
