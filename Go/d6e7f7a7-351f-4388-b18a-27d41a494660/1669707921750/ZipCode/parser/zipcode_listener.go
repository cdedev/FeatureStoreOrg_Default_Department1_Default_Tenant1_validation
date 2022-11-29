// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707921750/ZipCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ZipCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZipCodeListener is a complete listener for a parse tree produced by ZipCodeParser.
type ZipCodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterZipcode is called when entering the zipcode production.
	EnterZipcode(c *ZipcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitZipcode is called when exiting the zipcode production.
	ExitZipcode(c *ZipcodeContext)
}
