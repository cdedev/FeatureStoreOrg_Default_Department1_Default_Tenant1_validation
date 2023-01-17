// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924919318/Aquatic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aquatic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AquaticListener is a complete listener for a parse tree produced by AquaticParser.
type AquaticListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAquatic is called when entering the aquatic production.
	EnterAquatic(c *AquaticContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAquatic is called when exiting the aquatic production.
	ExitAquatic(c *AquaticContext)
}
