// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672379897809/Pain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PainListener is a complete listener for a parse tree produced by PainParser.
type PainListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPain is called when entering the pain production.
	EnterPain(c *PainContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPain is called when exiting the pain production.
	ExitPain(c *PainContext)
}
