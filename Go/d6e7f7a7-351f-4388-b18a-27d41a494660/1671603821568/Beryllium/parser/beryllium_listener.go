// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603821568/Beryllium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Beryllium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BerylliumListener is a complete listener for a parse tree produced by BerylliumParser.
type BerylliumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBeryllium is called when entering the beryllium production.
	EnterBeryllium(c *BerylliumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBeryllium is called when exiting the beryllium production.
	ExitBeryllium(c *BerylliumContext)
}
