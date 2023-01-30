// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072787036/Mortality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MortalityListener is a complete listener for a parse tree produced by MortalityParser.
type MortalityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMortality is called when entering the mortality production.
	EnterMortality(c *MortalityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMortality is called when exiting the mortality production.
	ExitMortality(c *MortalityContext)
}
