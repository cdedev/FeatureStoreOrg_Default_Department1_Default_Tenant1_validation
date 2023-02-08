// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836191886/DeathRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DeathRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DeathRateListener is a complete listener for a parse tree produced by DeathRateParser.
type DeathRateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDeathrate is called when entering the deathrate production.
	EnterDeathrate(c *DeathrateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDeathrate is called when exiting the deathrate production.
	ExitDeathrate(c *DeathrateContext)
}
