// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606868055/BloodGroup.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodGroup

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BloodGroupListener is a complete listener for a parse tree produced by BloodGroupParser.
type BloodGroupListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBloodgroup is called when entering the bloodgroup production.
	EnterBloodgroup(c *BloodgroupContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBloodgroup is called when exiting the bloodgroup production.
	ExitBloodgroup(c *BloodgroupContext)
}
