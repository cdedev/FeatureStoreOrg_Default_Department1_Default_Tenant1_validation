// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672978152440/ShipDate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ShipDate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShipDateListener is a complete listener for a parse tree produced by ShipDateParser.
type ShipDateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterShipdate is called when entering the shipdate production.
	EnterShipdate(c *ShipdateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitShipdate is called when exiting the shipdate production.
	ExitShipdate(c *ShipdateContext)
}
