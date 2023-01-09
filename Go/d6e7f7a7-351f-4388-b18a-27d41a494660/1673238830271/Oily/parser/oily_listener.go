// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238830271/Oily.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oily

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OilyListener is a complete listener for a parse tree produced by OilyParser.
type OilyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOily is called when entering the oily production.
	EnterOily(c *OilyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOily is called when exiting the oily production.
	ExitOily(c *OilyContext)
}
