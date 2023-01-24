// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532685649/Party.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Party

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePartyListener is a complete listener for a parse tree produced by PartyParser.
type BasePartyListener struct{}

var _ PartyListener = &BasePartyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePartyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePartyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePartyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePartyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePartyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePartyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterParty is called when production party is entered.
func (s *BasePartyListener) EnterParty(ctx *PartyContext) {}

// ExitParty is called when production party is exited.
func (s *BasePartyListener) ExitParty(ctx *PartyContext) {}
