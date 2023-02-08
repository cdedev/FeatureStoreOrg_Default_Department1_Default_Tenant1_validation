// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834231475/Members.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Members

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MembersListener is a complete listener for a parse tree produced by MembersParser.
type MembersListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMembers is called when entering the members production.
	EnterMembers(c *MembersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMembers is called when exiting the members production.
	ExitMembers(c *MembersContext)
}
