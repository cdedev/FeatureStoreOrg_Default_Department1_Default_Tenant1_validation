// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672117233276/MarketCap.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarketCap

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarketCapListener is a complete listener for a parse tree produced by MarketCapParser.
type MarketCapListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMarketcap is called when entering the marketcap production.
	EnterMarketcap(c *MarketcapContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMarketcap is called when exiting the marketcap production.
	ExitMarketcap(c *MarketcapContext)
}
