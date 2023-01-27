// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674790494867/Mass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MassListener is a complete listener for a parse tree produced by MassParser.
type MassListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMass is called when entering the mass production.
	EnterMass(c *MassContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMass is called when exiting the mass production.
	ExitMass(c *MassContext)
}
