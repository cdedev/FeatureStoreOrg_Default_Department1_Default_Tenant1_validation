// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672631742428/Term.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Term

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TermListener is a complete listener for a parse tree produced by TermParser.
type TermListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
