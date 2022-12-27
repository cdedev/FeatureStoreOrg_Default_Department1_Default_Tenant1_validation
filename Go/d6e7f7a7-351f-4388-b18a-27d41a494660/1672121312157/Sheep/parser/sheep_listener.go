// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121312157/Sheep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sheep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SheepListener is a complete listener for a parse tree produced by SheepParser.
type SheepListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSheep is called when entering the sheep production.
	EnterSheep(c *SheepContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSheep is called when exiting the sheep production.
	ExitSheep(c *SheepContext)
}
