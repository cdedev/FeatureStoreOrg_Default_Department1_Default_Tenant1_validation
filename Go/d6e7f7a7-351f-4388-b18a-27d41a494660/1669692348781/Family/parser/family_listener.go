// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692348781/Family.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Family

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FamilyListener is a complete listener for a parse tree produced by FamilyParser.
type FamilyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFamily is called when entering the family production.
	EnterFamily(c *FamilyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFamily is called when exiting the family production.
	ExitFamily(c *FamilyContext)
}
