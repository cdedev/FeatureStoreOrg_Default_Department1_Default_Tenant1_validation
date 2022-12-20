// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671525329885/Population.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Population

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PopulationListener is a complete listener for a parse tree produced by PopulationParser.
type PopulationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPopulation is called when entering the population production.
	EnterPopulation(c *PopulationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPopulation is called when exiting the population production.
	ExitPopulation(c *PopulationContext)
}
