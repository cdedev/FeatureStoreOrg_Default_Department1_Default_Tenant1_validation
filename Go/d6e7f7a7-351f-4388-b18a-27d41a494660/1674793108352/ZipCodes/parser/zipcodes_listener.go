// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793108352/ZipCodes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ZipCodes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZipCodesListener is a complete listener for a parse tree produced by ZipCodesParser.
type ZipCodesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterZipcodes is called when entering the zipcodes production.
	EnterZipcodes(c *ZipcodesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitZipcodes is called when exiting the zipcodes production.
	ExitZipcodes(c *ZipcodesContext)
}
