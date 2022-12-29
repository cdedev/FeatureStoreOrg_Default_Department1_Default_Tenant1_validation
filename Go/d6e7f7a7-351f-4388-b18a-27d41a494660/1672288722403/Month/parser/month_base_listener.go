// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288722403/Month.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Month

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMonthListener is a complete listener for a parse tree produced by MonthParser.
type BaseMonthListener struct{}

var _ MonthListener = &BaseMonthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMonthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMonthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMonthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMonthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMonthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMonthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseMonthListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseMonthListener) ExitMonth(ctx *MonthContext) {}
