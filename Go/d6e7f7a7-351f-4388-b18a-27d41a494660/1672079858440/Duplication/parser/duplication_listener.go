// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672079858440/Duplication.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Duplication

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DuplicationListener is a complete listener for a parse tree produced by DuplicationParser.
type DuplicationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDuplication is called when entering the duplication production.
	EnterDuplication(c *DuplicationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDuplication is called when exiting the duplication production.
	ExitDuplication(c *DuplicationContext)
}
