// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675074640883/DateOfBirth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfBirth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DateOfBirthListener is a complete listener for a parse tree produced by DateOfBirthParser.
type DateOfBirthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDateofbirth is called when entering the dateofbirth production.
	EnterDateofbirth(c *DateofbirthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDateofbirth is called when exiting the dateofbirth production.
	ExitDateofbirth(c *DateofbirthContext)
}
