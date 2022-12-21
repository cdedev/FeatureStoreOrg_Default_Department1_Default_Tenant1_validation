// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671599159281/Entity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EntityListener is a complete listener for a parse tree produced by EntityParser.
type EntityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEntity is called when entering the entity production.
	EnterEntity(c *EntityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEntity is called when exiting the entity production.
	ExitEntity(c *EntityContext)
}
