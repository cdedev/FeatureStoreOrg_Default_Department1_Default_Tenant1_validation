// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671514059387/Ownership.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ownership

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OwnershipListener is a complete listener for a parse tree produced by OwnershipParser.
type OwnershipListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOwnership is called when entering the ownership production.
	EnterOwnership(c *OwnershipContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOwnership is called when exiting the ownership production.
	ExitOwnership(c *OwnershipContext)
}
