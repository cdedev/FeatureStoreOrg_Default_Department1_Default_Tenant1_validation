// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669625336662/Strength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Strength

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StrengthListener is a complete listener for a parse tree produced by StrengthParser.
type StrengthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStrength is called when entering the strength production.
	EnterStrength(c *StrengthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStrength is called when exiting the strength production.
	ExitStrength(c *StrengthContext)
}
