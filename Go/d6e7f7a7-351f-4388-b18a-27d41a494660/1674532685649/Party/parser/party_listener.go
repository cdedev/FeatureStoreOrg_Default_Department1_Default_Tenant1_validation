// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532685649/Party.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Party

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PartyListener is a complete listener for a parse tree produced by PartyParser.
type PartyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterParty is called when entering the party production.
	EnterParty(c *PartyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitParty is called when exiting the party production.
	ExitParty(c *PartyContext)
}
