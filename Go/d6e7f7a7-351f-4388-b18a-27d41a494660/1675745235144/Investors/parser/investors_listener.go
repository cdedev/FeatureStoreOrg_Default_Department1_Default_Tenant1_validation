// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675745235144/Investors.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Investors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InvestorsListener is a complete listener for a parse tree produced by InvestorsParser.
type InvestorsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInvestors is called when entering the investors production.
	EnterInvestors(c *InvestorsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInvestors is called when exiting the investors production.
	ExitInvestors(c *InvestorsContext)
}
