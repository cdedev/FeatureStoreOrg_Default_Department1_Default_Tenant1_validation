// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698569487/YearsInCurrentRole.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // YearsInCurrentRole

import "github.com/antlr/antlr4/runtime/Go/antlr"

// YearsInCurrentRoleListener is a complete listener for a parse tree produced by YearsInCurrentRoleParser.
type YearsInCurrentRoleListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterYearsincurrentrole is called when entering the yearsincurrentrole production.
	EnterYearsincurrentrole(c *YearsincurrentroleContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitYearsincurrentrole is called when exiting the yearsincurrentrole production.
	ExitYearsincurrentrole(c *YearsincurrentroleContext)
}
