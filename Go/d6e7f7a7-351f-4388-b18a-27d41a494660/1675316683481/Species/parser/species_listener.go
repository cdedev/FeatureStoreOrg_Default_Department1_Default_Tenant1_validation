// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675316683481/Species.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Species

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpeciesListener is a complete listener for a parse tree produced by SpeciesParser.
type SpeciesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSpecies is called when entering the species production.
	EnterSpecies(c *SpeciesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSpecies is called when exiting the species production.
	ExitSpecies(c *SpeciesContext)
}
