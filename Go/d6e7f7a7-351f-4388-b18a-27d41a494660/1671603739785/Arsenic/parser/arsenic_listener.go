// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603739785/Arsenic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Arsenic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ArsenicListener is a complete listener for a parse tree produced by ArsenicParser.
type ArsenicListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArsenic is called when entering the arsenic production.
	EnterArsenic(c *ArsenicContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArsenic is called when exiting the arsenic production.
	ExitArsenic(c *ArsenicContext)
}
