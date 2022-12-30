// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672375214175/FamilySize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FamilySize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FamilySizeListener is a complete listener for a parse tree produced by FamilySizeParser.
type FamilySizeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFamilysize is called when entering the familysize production.
	EnterFamilysize(c *FamilysizeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFamilysize is called when exiting the familysize production.
	ExitFamilysize(c *FamilysizeContext)
}
