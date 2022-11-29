// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697837214/Chlorides.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chlorides

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChloridesListener is a complete listener for a parse tree produced by ChloridesParser.
type ChloridesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChlorides is called when entering the chlorides production.
	EnterChlorides(c *ChloridesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChlorides is called when exiting the chlorides production.
	ExitChlorides(c *ChloridesContext)
}
