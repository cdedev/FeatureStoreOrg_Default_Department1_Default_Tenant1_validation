// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516017000/BranchCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BranchCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BranchCodeListener is a complete listener for a parse tree produced by BranchCodeParser.
type BranchCodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBranchcode is called when entering the branchcode production.
	EnterBranchcode(c *BranchcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBranchcode is called when exiting the branchcode production.
	ExitBranchcode(c *BranchcodeContext)
}
