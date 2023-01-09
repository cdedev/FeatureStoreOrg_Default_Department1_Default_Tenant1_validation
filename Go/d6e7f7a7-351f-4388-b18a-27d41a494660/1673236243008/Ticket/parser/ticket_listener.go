// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236243008/Ticket.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ticket

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TicketListener is a complete listener for a parse tree produced by TicketParser.
type TicketListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTicket is called when entering the ticket production.
	EnterTicket(c *TicketContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTicket is called when exiting the ticket production.
	ExitTicket(c *TicketContext)
}
