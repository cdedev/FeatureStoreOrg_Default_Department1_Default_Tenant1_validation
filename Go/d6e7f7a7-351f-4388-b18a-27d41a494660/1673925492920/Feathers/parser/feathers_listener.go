// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925492920/Feathers.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Feathers

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FeathersListener is a complete listener for a parse tree produced by FeathersParser.
type FeathersListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFeathers is called when entering the feathers production.
	EnterFeathers(c *FeathersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFeathers is called when exiting the feathers production.
	ExitFeathers(c *FeathersContext)
}
