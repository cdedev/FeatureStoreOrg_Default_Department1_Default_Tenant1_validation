// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674623607607/Stock.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stock

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StockListener is a complete listener for a parse tree produced by StockParser.
type StockListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStock is called when entering the stock production.
	EnterStock(c *StockContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStock is called when exiting the stock production.
	ExitStock(c *StockContext)
}
