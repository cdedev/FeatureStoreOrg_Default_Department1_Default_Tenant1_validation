// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884745703/Town.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Town

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TownListener is a complete listener for a parse tree produced by TownParser.
type TownListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTown is called when entering the town production.
	EnterTown(c *TownContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTown is called when exiting the town production.
	ExitTown(c *TownContext)
}
