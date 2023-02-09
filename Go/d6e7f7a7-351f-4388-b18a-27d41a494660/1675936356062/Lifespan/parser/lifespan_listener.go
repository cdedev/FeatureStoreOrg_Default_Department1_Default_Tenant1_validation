// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675936356062/Lifespan.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lifespan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LifespanListener is a complete listener for a parse tree produced by LifespanParser.
type LifespanListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLifespan is called when entering the lifespan production.
	EnterLifespan(c *LifespanContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLifespan is called when exiting the lifespan production.
	ExitLifespan(c *LifespanContext)
}
