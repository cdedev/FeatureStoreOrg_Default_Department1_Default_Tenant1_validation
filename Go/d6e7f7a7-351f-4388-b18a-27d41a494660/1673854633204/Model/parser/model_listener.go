// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673854633204/Model.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Model

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ModelListener is a complete listener for a parse tree produced by ModelParser.
type ModelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterModel is called when entering the model production.
	EnterModel(c *ModelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitModel is called when exiting the model production.
	ExitModel(c *ModelContext)
}
