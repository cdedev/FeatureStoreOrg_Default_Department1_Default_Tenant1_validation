// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672200041640/House.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // House

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HouseListener is a complete listener for a parse tree produced by HouseParser.
type HouseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHouse is called when entering the house production.
	EnterHouse(c *HouseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHouse is called when exiting the house production.
	ExitHouse(c *HouseContext)
}
