// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669786719095/Market.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Market

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarketListener is a complete listener for a parse tree produced by MarketParser.
type MarketListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMarket is called when entering the market production.
	EnterMarket(c *MarketContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMarket is called when exiting the market production.
	ExitMarket(c *MarketContext)
}
