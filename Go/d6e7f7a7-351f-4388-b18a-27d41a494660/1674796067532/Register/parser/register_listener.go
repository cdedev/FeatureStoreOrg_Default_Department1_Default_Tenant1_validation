// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674796067532/Register.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Register

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RegisterListener is a complete listener for a parse tree produced by RegisterParser.
type RegisterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)
}
