// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676133282849/Bills.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bills

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BillsListener is a complete listener for a parse tree produced by BillsParser.
type BillsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBills is called when entering the bills production.
	EnterBills(c *BillsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBills is called when exiting the bills production.
	ExitBills(c *BillsContext)
}
