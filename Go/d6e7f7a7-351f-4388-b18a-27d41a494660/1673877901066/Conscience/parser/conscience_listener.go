// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673877901066/Conscience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Conscience

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConscienceListener is a complete listener for a parse tree produced by ConscienceParser.
type ConscienceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConscience is called when entering the conscience production.
	EnterConscience(c *ConscienceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConscience is called when exiting the conscience production.
	ExitConscience(c *ConscienceContext)
}
