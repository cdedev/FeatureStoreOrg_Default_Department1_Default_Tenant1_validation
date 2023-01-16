// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878269550/Paresthesia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Paresthesia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ParesthesiaListener is a complete listener for a parse tree produced by ParesthesiaParser.
type ParesthesiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterParesthesia is called when entering the paresthesia production.
	EnterParesthesia(c *ParesthesiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitParesthesia is called when exiting the paresthesia production.
	ExitParesthesia(c *ParesthesiaContext)
}
