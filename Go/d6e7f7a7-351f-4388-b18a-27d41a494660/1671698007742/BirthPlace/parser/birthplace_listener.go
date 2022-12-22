// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698007742/BirthPlace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BirthPlace

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BirthPlaceListener is a complete listener for a parse tree produced by BirthPlaceParser.
type BirthPlaceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBirthplace is called when entering the birthplace production.
	EnterBirthplace(c *BirthplaceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBirthplace is called when exiting the birthplace production.
	ExitBirthplace(c *BirthplaceContext)
}
