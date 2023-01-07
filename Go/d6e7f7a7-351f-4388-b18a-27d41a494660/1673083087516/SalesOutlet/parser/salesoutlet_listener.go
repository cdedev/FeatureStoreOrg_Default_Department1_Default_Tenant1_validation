// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083087516/SalesOutlet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SalesOutlet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SalesOutletListener is a complete listener for a parse tree produced by SalesOutletParser.
type SalesOutletListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSalesoutlet is called when entering the salesoutlet production.
	EnterSalesoutlet(c *SalesoutletContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSalesoutlet is called when exiting the salesoutlet production.
	ExitSalesoutlet(c *SalesoutletContext)
}
