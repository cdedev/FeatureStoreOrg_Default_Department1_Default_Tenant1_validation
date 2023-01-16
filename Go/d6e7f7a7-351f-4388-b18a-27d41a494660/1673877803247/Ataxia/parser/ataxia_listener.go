// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673877803247/Ataxia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ataxia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AtaxiaListener is a complete listener for a parse tree produced by AtaxiaParser.
type AtaxiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAtaxia is called when entering the ataxia production.
	EnterAtaxia(c *AtaxiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAtaxia is called when exiting the ataxia production.
	ExitAtaxia(c *AtaxiaContext)
}
