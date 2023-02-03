// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675392826128/Concentration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concentration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConcentrationListener is a complete listener for a parse tree produced by ConcentrationParser.
type ConcentrationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConcentration is called when entering the concentration production.
	EnterConcentration(c *ConcentrationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConcentration is called when exiting the concentration production.
	ExitConcentration(c *ConcentrationContext)
}
