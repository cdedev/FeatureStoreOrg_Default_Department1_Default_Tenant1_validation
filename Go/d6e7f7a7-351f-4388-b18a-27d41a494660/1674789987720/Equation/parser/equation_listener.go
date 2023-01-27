// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789987720/Equation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Equation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EquationListener is a complete listener for a parse tree produced by EquationParser.
type EquationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)
}
