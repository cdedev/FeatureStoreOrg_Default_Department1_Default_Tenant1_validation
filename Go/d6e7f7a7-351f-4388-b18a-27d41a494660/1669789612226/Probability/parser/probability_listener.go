// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789612226/Probability.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Probability

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProbabilityListener is a complete listener for a parse tree produced by ProbabilityParser.
type ProbabilityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProbability is called when entering the probability production.
	EnterProbability(c *ProbabilityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProbability is called when exiting the probability production.
	ExitProbability(c *ProbabilityContext)
}
