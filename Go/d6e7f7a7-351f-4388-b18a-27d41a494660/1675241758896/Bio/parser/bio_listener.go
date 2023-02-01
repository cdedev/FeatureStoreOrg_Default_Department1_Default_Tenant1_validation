// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241758896/Bio.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bio

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BioListener is a complete listener for a parse tree produced by BioParser.
type BioListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBio is called when entering the bio production.
	EnterBio(c *BioContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBio is called when exiting the bio production.
	ExitBio(c *BioContext)
}
