// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603911590/Bohrium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bohrium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BohriumListener is a complete listener for a parse tree produced by BohriumParser.
type BohriumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBohrium is called when entering the bohrium production.
	EnterBohrium(c *BohriumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBohrium is called when exiting the bohrium production.
	ExitBohrium(c *BohriumContext)
}
