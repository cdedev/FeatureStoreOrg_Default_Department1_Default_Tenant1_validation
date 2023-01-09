// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238650276/Combination.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Combination

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CombinationListener is a complete listener for a parse tree produced by CombinationParser.
type CombinationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCombination is called when entering the combination production.
	EnterCombination(c *CombinationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCombination is called when exiting the combination production.
	ExitCombination(c *CombinationContext)
}
