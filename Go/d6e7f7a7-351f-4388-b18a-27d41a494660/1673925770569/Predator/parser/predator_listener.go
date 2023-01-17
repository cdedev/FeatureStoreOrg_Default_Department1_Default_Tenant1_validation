// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925770569/Predator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Predator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PredatorListener is a complete listener for a parse tree produced by PredatorParser.
type PredatorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPredator is called when entering the predator production.
	EnterPredator(c *PredatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPredator is called when exiting the predator production.
	ExitPredator(c *PredatorContext)
}
