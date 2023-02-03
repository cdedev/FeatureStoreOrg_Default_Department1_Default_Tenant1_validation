// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675397755196/BloodSodium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodSodium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BloodSodiumListener is a complete listener for a parse tree produced by BloodSodiumParser.
type BloodSodiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBloodsodium is called when entering the bloodsodium production.
	EnterBloodsodium(c *BloodsodiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBloodsodium is called when exiting the bloodsodium production.
	ExitBloodsodium(c *BloodsodiumContext)
}
