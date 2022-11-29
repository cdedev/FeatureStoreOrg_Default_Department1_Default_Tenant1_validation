// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690493750/Sodium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sodium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SodiumListener is a complete listener for a parse tree produced by SodiumParser.
type SodiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSodium is called when entering the sodium production.
	EnterSodium(c *SodiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSodium is called when exiting the sodium production.
	ExitSodium(c *SodiumContext)
}
