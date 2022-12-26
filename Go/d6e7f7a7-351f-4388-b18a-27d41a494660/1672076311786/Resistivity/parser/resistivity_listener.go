// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076311786/Resistivity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Resistivity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ResistivityListener is a complete listener for a parse tree produced by ResistivityParser.
type ResistivityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterResistivity is called when entering the resistivity production.
	EnterResistivity(c *ResistivityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitResistivity is called when exiting the resistivity production.
	ExitResistivity(c *ResistivityContext)
}
