// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672204141055/FederalOccupation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FederalOccupation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FederalOccupationListener is a complete listener for a parse tree produced by FederalOccupationParser.
type FederalOccupationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFederaloccupation is called when entering the federaloccupation production.
	EnterFederaloccupation(c *FederaloccupationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFederaloccupation is called when exiting the federaloccupation production.
	ExitFederaloccupation(c *FederaloccupationContext)
}
