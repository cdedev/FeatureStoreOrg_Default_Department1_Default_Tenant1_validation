// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672977218378/WeekdayName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeekdayName

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeekdayNameListener is a complete listener for a parse tree produced by WeekdayNameParser.
type BaseWeekdayNameListener struct{}

var _ WeekdayNameListener = &BaseWeekdayNameListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeekdayNameListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeekdayNameListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeekdayNameListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeekdayNameListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeekdayNameListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeekdayNameListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeekdayname is called when production weekdayname is entered.
func (s *BaseWeekdayNameListener) EnterWeekdayname(ctx *WeekdaynameContext) {}

// ExitWeekdayname is called when production weekdayname is exited.
func (s *BaseWeekdayNameListener) ExitWeekdayname(ctx *WeekdaynameContext) {}
