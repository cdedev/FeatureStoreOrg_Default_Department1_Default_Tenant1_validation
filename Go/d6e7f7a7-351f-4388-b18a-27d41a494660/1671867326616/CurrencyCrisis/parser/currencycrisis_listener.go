// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867326616/CurrencyCrisis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CurrencyCrisis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CurrencyCrisisListener is a complete listener for a parse tree produced by CurrencyCrisisParser.
type CurrencyCrisisListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCurrencycrisis is called when entering the currencycrisis production.
	EnterCurrencycrisis(c *CurrencycrisisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCurrencycrisis is called when exiting the currencycrisis production.
	ExitCurrencycrisis(c *CurrencycrisisContext)
}
