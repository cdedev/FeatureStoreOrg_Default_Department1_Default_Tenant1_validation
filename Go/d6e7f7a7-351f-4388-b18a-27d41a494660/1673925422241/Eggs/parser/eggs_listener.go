// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925422241/Eggs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eggs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EggsListener is a complete listener for a parse tree produced by EggsParser.
type EggsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEggs is called when entering the eggs production.
	EnterEggs(c *EggsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEggs is called when exiting the eggs production.
	ExitEggs(c *EggsContext)
}
