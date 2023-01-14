// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673672257964/Currency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Currency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CurrencyListener is a complete listener for a parse tree produced by CurrencyParser.
type CurrencyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCurrency is called when entering the currency production.
	EnterCurrency(c *CurrencyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCurrency is called when exiting the currency production.
	ExitCurrency(c *CurrencyContext)
}
