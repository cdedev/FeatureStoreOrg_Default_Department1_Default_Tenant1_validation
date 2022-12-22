// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671699291388/Infection.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Infection

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InfectionListener is a complete listener for a parse tree produced by InfectionParser.
type InfectionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInfection is called when entering the infection production.
	EnterInfection(c *InfectionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInfection is called when exiting the infection production.
	ExitInfection(c *InfectionContext)
}
