// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604371088/Cerium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cerium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CeriumListener is a complete listener for a parse tree produced by CeriumParser.
type CeriumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCerium is called when entering the cerium production.
	EnterCerium(c *CeriumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCerium is called when exiting the cerium production.
	ExitCerium(c *CeriumContext)
}
