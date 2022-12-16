// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206491780/Infected.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Infected

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InfectedListener is a complete listener for a parse tree produced by InfectedParser.
type InfectedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInfected is called when entering the infected production.
	EnterInfected(c *InfectedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInfected is called when exiting the infected production.
	ExitInfected(c *InfectedContext)
}
