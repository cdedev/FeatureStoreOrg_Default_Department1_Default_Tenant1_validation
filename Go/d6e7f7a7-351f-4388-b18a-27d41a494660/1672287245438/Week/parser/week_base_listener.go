// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672287245438/Week.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Week

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeekListener is a complete listener for a parse tree produced by WeekParser.
type BaseWeekListener struct{}

var _ WeekListener = &BaseWeekListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeekListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeekListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeekListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeekListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeekListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeekListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeek is called when production week is entered.
func (s *BaseWeekListener) EnterWeek(ctx *WeekContext) {}

// ExitWeek is called when production week is exited.
func (s *BaseWeekListener) ExitWeek(ctx *WeekContext) {}
