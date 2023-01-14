// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676512220/Symmetry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Symmetry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SymmetryListener is a complete listener for a parse tree produced by SymmetryParser.
type SymmetryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSymmetry is called when entering the symmetry production.
	EnterSymmetry(c *SymmetryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSymmetry is called when exiting the symmetry production.
	ExitSymmetry(c *SymmetryContext)
}
