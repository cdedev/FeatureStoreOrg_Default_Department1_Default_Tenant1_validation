// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878520926/Tinnitus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinnitus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TinnitusListener is a complete listener for a parse tree produced by TinnitusParser.
type TinnitusListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTinnitus is called when entering the tinnitus production.
	EnterTinnitus(c *TinnitusContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTinnitus is called when exiting the tinnitus production.
	ExitTinnitus(c *TinnitusContext)
}
