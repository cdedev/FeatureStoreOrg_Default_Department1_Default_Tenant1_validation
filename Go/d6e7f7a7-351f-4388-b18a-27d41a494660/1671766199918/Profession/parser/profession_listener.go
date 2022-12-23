// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671766199918/Profession.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Profession

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProfessionListener is a complete listener for a parse tree produced by ProfessionParser.
type ProfessionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProfession is called when entering the profession production.
	EnterProfession(c *ProfessionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProfession is called when exiting the profession production.
	ExitProfession(c *ProfessionContext)
}
