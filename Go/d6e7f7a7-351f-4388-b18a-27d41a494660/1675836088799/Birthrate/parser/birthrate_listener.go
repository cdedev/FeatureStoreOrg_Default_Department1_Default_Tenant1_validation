// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836088799/Birthrate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Birthrate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BirthrateListener is a complete listener for a parse tree produced by BirthrateParser.
type BirthrateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBirthrate is called when entering the birthrate production.
	EnterBirthrate(c *BirthrateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBirthrate is called when exiting the birthrate production.
	ExitBirthrate(c *BirthrateContext)
}
