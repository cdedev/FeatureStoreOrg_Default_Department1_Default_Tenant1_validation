// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669981590412/Date_of_birth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Date_of_birth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Date_of_birthListener is a complete listener for a parse tree produced by Date_of_birthParser.
type Date_of_birthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDate_of_birth is called when entering the date_of_birth production.
	EnterDate_of_birth(c *Date_of_birthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDate_of_birth is called when exiting the date_of_birth production.
	ExitDate_of_birth(c *Date_of_birthContext)
}
