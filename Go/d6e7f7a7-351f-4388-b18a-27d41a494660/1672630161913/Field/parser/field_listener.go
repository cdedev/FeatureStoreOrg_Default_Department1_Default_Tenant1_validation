// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672630161913/Field.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Field

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FieldListener is a complete listener for a parse tree produced by FieldParser.
type FieldListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)
}
