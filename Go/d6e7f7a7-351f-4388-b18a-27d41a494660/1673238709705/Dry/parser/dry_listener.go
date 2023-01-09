// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238709705/Dry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DryListener is a complete listener for a parse tree produced by DryParser.
type DryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDry is called when entering the dry production.
	EnterDry(c *DryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDry is called when exiting the dry production.
	ExitDry(c *DryContext)
}
