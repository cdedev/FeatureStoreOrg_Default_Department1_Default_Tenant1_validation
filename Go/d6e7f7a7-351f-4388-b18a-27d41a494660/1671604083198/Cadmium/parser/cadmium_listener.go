// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604083198/Cadmium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cadmium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CadmiumListener is a complete listener for a parse tree produced by CadmiumParser.
type CadmiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCadmium is called when entering the cadmium production.
	EnterCadmium(c *CadmiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCadmium is called when exiting the cadmium production.
	ExitCadmium(c *CadmiumContext)
}
