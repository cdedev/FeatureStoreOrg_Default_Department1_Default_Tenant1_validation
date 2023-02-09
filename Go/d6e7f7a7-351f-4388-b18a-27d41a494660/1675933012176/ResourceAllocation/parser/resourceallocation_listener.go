// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675933012176/ResourceAllocation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ResourceAllocation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ResourceAllocationListener is a complete listener for a parse tree produced by ResourceAllocationParser.
type ResourceAllocationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterResourceallocation is called when entering the resourceallocation production.
	EnterResourceallocation(c *ResourceallocationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitResourceallocation is called when exiting the resourceallocation production.
	ExitResourceallocation(c *ResourceallocationContext)
}
