// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626974586/Gender.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gender

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GenderListener is a complete listener for a parse tree produced by GenderParser.
type GenderListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGender is called when entering the gender production.
	EnterGender(c *GenderContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGender is called when exiting the gender production.
	ExitGender(c *GenderContext)
}
