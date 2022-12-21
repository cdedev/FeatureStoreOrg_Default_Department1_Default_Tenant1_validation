// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603781202/Barium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BariumListener is a complete listener for a parse tree produced by BariumParser.
type BariumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBarium is called when entering the barium production.
	EnterBarium(c *BariumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBarium is called when exiting the barium production.
	ExitBarium(c *BariumContext)
}
