// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528680321/Residence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Residence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ResidenceListener is a complete listener for a parse tree produced by ResidenceParser.
type ResidenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterResidence is called when entering the residence production.
	EnterResidence(c *ResidenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitResidence is called when exiting the residence production.
	ExitResidence(c *ResidenceContext)
}
