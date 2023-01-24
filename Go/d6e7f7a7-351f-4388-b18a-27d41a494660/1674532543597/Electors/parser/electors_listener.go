// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532543597/Electors.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Electors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ElectorsListener is a complete listener for a parse tree produced by ElectorsParser.
type ElectorsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterElectors is called when entering the electors production.
	EnterElectors(c *ElectorsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitElectors is called when exiting the electors production.
	ExitElectors(c *ElectorsContext)
}
