// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674534363476/MinorityVote.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MinorityVote

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MinorityVoteListener is a complete listener for a parse tree produced by MinorityVoteParser.
type MinorityVoteListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMinorityvote is called when entering the minorityvote production.
	EnterMinorityvote(c *MinorityvoteContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMinorityvote is called when exiting the minorityvote production.
	ExitMinorityvote(c *MinorityvoteContext)
}
