// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672636987589/Entry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEntryListener is a complete listener for a parse tree produced by EntryParser.
type BaseEntryListener struct{}

var _ EntryListener = &BaseEntryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEntryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEntryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEntryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEntryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEntryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEntryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseEntryListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseEntryListener) ExitEntry(ctx *EntryContext) {}
