// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672380066598/Fever.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fever

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FeverListener is a complete listener for a parse tree produced by FeverParser.
type FeverListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFever is called when entering the fever production.
	EnterFever(c *FeverContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFever is called when exiting the fever production.
	ExitFever(c *FeverContext)
}
