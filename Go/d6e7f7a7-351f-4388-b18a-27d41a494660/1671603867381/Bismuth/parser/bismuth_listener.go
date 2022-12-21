// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603867381/Bismuth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bismuth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BismuthListener is a complete listener for a parse tree produced by BismuthParser.
type BismuthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBismuth is called when entering the bismuth production.
	EnterBismuth(c *BismuthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBismuth is called when exiting the bismuth production.
	ExitBismuth(c *BismuthContext)
}
