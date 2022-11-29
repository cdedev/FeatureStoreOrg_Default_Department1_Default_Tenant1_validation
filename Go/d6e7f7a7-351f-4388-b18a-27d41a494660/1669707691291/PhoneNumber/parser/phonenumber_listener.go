// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707691291/PhoneNumber.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneNumber

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhoneNumberListener is a complete listener for a parse tree produced by PhoneNumberParser.
type PhoneNumberListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPhonenumber is called when entering the phonenumber production.
	EnterPhonenumber(c *PhonenumberContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPhonenumber is called when exiting the phonenumber production.
	ExitPhonenumber(c *PhonenumberContext)
}
