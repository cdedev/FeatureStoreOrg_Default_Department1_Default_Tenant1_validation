// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672722988491/Nationality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nationality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NationalityListener is a complete listener for a parse tree produced by NationalityParser.
type NationalityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNationality is called when entering the nationality production.
	EnterNationality(c *NationalityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNationality is called when exiting the nationality production.
	ExitNationality(c *NationalityContext)
}
