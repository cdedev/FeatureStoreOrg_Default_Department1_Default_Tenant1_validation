// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883355662/Event.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Event

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EventListener is a complete listener for a parse tree produced by EventParser.
type EventListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)
}
