// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671597234897/Stress.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stress

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StressListener is a complete listener for a parse tree produced by StressParser.
type StressListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStress is called when entering the stress production.
	EnterStress(c *StressContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStress is called when exiting the stress production.
	ExitStress(c *StressContext)
}
