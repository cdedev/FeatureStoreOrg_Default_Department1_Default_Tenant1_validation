// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878575030/Vertigo.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vertigo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VertigoListener is a complete listener for a parse tree produced by VertigoParser.
type VertigoListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVertigo is called when entering the vertigo production.
	EnterVertigo(c *VertigoContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVertigo is called when exiting the vertigo production.
	ExitVertigo(c *VertigoContext)
}
