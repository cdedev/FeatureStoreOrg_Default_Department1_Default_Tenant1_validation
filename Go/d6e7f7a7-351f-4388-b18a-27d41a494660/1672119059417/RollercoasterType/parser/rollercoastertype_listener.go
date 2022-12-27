// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119059417/RollercoasterType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RollercoasterType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RollercoasterTypeListener is a complete listener for a parse tree produced by RollercoasterTypeParser.
type RollercoasterTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRollercoastertype is called when entering the rollercoastertype production.
	EnterRollercoastertype(c *RollercoastertypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRollercoastertype is called when exiting the rollercoastertype production.
	ExitRollercoastertype(c *RollercoastertypeContext)
}
