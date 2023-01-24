// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532749322/Votes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Votes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VotesListener is a complete listener for a parse tree produced by VotesParser.
type VotesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVotes is called when entering the votes production.
	EnterVotes(c *VotesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVotes is called when exiting the votes production.
	ExitVotes(c *VotesContext)
}
