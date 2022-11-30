// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778760999/PhoneService.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneService

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhoneServiceListener is a complete listener for a parse tree produced by PhoneServiceParser.
type PhoneServiceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPhoneservice is called when entering the phoneservice production.
	EnterPhoneservice(c *PhoneserviceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPhoneservice is called when exiting the phoneservice production.
	ExitPhoneservice(c *PhoneserviceContext)
}
