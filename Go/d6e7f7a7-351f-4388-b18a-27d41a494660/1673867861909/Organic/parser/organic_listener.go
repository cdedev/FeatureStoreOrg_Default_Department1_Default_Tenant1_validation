// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867861909/Organic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Organic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OrganicListener is a complete listener for a parse tree produced by OrganicParser.
type OrganicListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrganic is called when entering the organic production.
	EnterOrganic(c *OrganicContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrganic is called when exiting the organic production.
	ExitOrganic(c *OrganicContext)
}
