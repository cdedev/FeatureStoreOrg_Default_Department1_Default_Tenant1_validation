// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671530275929/Commodity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Commodity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CommodityListener is a complete listener for a parse tree produced by CommodityParser.
type CommodityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCommodity is called when entering the commodity production.
	EnterCommodity(c *CommodityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCommodity is called when exiting the commodity production.
	ExitCommodity(c *CommodityContext)
}
