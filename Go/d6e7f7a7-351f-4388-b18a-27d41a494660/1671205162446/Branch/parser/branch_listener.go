// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671205162446/Branch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Branch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BranchListener is a complete listener for a parse tree produced by BranchParser.
type BranchListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)
}
