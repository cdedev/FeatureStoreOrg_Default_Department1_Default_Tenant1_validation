// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674097827714/Registration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Registration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RegistrationListener is a complete listener for a parse tree produced by RegistrationParser.
type RegistrationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRegistration is called when entering the registration production.
	EnterRegistration(c *RegistrationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRegistration is called when exiting the registration production.
	ExitRegistration(c *RegistrationContext)
}
