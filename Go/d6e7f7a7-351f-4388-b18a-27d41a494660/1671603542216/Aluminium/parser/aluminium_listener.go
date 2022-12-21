// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603542216/Aluminium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aluminium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AluminiumListener is a complete listener for a parse tree produced by AluminiumParser.
type AluminiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAluminium is called when entering the aluminium production.
	EnterAluminium(c *AluminiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAluminium is called when exiting the aluminium production.
	ExitAluminium(c *AluminiumContext)
}
