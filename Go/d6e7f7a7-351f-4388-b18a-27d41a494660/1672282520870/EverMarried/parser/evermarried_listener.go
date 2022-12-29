// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672282520870/EverMarried.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EverMarried

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EverMarriedListener is a complete listener for a parse tree produced by EverMarriedParser.
type EverMarriedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEvermarried is called when entering the evermarried production.
	EnterEvermarried(c *EvermarriedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEvermarried is called when exiting the evermarried production.
	ExitEvermarried(c *EvermarriedContext)
}
