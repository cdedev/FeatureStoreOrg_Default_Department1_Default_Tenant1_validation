// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236243008/Ticket.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ticket

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTicketListener is a complete listener for a parse tree produced by TicketParser.
type BaseTicketListener struct{}

var _ TicketListener = &BaseTicketListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTicketListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTicketListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTicketListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTicketListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTicketListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTicketListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTicket is called when production ticket is entered.
func (s *BaseTicketListener) EnterTicket(ctx *TicketContext) {}

// ExitTicket is called when production ticket is exited.
func (s *BaseTicketListener) ExitTicket(ctx *TicketContext) {}
