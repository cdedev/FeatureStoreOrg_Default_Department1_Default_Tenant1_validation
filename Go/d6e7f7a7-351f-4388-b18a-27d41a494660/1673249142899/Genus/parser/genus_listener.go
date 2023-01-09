// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673249142899/Genus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Genus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GenusListener is a complete listener for a parse tree produced by GenusParser.
type GenusListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGenus is called when entering the genus production.
	EnterGenus(c *GenusContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGenus is called when exiting the genus production.
	ExitGenus(c *GenusContext)
}
