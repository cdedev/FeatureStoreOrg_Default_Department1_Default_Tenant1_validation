// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675317813147/Attribute.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Attribute

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AttributeListener is a complete listener for a parse tree produced by AttributeParser.
type AttributeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)
}
