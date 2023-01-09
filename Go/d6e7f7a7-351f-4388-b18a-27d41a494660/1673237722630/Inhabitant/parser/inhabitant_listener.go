// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237722630/Inhabitant.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Inhabitant

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InhabitantListener is a complete listener for a parse tree produced by InhabitantParser.
type InhabitantListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInhabitant is called when entering the inhabitant production.
	EnterInhabitant(c *InhabitantContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInhabitant is called when exiting the inhabitant production.
	ExitInhabitant(c *InhabitantContext)
}
