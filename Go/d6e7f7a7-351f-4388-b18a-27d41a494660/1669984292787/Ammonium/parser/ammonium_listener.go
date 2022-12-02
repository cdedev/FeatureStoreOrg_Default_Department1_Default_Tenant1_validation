// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669984292787/Ammonium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ammonium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AmmoniumListener is a complete listener for a parse tree produced by AmmoniumParser.
type AmmoniumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAmmonium is called when entering the ammonium production.
	EnterAmmonium(c *AmmoniumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAmmonium is called when exiting the ammonium production.
	ExitAmmonium(c *AmmoniumContext)
}
