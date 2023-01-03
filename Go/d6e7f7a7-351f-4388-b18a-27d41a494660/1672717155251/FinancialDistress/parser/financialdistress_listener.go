// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717155251/FinancialDistress.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FinancialDistress

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FinancialDistressListener is a complete listener for a parse tree produced by FinancialDistressParser.
type FinancialDistressListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFinancialdistress is called when entering the financialdistress production.
	EnterFinancialdistress(c *FinancialdistressContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFinancialdistress is called when exiting the financialdistress production.
	ExitFinancialdistress(c *FinancialdistressContext)
}
