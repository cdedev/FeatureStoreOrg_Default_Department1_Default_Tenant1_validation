// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671608397452/Ram.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ram

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RamListener is a complete listener for a parse tree produced by RamParser.
type RamListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRam is called when entering the ram production.
	EnterRam(c *RamContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRam is called when exiting the ram production.
	ExitRam(c *RamContext)
}
