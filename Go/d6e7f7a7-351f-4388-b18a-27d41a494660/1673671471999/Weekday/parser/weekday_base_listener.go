// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673671471999/Weekday.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weekday

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeekdayListener is a complete listener for a parse tree produced by WeekdayParser.
type BaseWeekdayListener struct{}

var _ WeekdayListener = &BaseWeekdayListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeekdayListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeekdayListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeekdayListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeekdayListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeekdayListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeekdayListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeekday is called when production weekday is entered.
func (s *BaseWeekdayListener) EnterWeekday(ctx *WeekdayContext) {}

// ExitWeekday is called when production weekday is exited.
func (s *BaseWeekdayListener) ExitWeekday(ctx *WeekdayContext) {}
