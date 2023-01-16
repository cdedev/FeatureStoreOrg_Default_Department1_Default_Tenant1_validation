// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878039204/Diplopia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diplopia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiplopiaListener is a complete listener for a parse tree produced by DiplopiaParser.
type DiplopiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiplopia is called when entering the diplopia production.
	EnterDiplopia(c *DiplopiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiplopia is called when exiting the diplopia production.
	ExitDiplopia(c *DiplopiaContext)
}
