// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674791022156/Substance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Substance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SubstanceListener is a complete listener for a parse tree produced by SubstanceParser.
type SubstanceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubstance is called when entering the substance production.
	EnterSubstance(c *SubstanceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubstance is called when exiting the substance production.
	ExitSubstance(c *SubstanceContext)
}
