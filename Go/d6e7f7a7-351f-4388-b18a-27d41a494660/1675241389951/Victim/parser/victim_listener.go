// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241389951/Victim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Victim

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VictimListener is a complete listener for a parse tree produced by VictimParser.
type VictimListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVictim is called when entering the victim production.
	EnterVictim(c *VictimContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVictim is called when exiting the victim production.
	ExitVictim(c *VictimContext)
}
