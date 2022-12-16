// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206430189/Deathcases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Deathcases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DeathcasesListener is a complete listener for a parse tree produced by DeathcasesParser.
type DeathcasesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDeathcases is called when entering the deathcases production.
	EnterDeathcases(c *DeathcasesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDeathcases is called when exiting the deathcases production.
	ExitDeathcases(c *DeathcasesContext)
}
