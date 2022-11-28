// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655720318/BloodOxygen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodOxygen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BloodOxygenListener is a complete listener for a parse tree produced by BloodOxygenParser.
type BloodOxygenListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBloodoxygen is called when entering the bloodoxygen production.
	EnterBloodoxygen(c *BloodoxygenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBloodoxygen is called when exiting the bloodoxygen production.
	ExitBloodoxygen(c *BloodoxygenContext)
}
