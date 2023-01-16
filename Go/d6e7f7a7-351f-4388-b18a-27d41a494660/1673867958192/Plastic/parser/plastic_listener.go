// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867958192/Plastic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Plastic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PlasticListener is a complete listener for a parse tree produced by PlasticParser.
type PlasticListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPlastic is called when entering the plastic production.
	EnterPlastic(c *PlasticContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPlastic is called when exiting the plastic production.
	ExitPlastic(c *PlasticContext)
}
