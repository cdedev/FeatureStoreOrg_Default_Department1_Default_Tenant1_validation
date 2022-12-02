// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669974295055/Holiday.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Holiday

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHolidayListener is a complete listener for a parse tree produced by HolidayParser.
type BaseHolidayListener struct{}

var _ HolidayListener = &BaseHolidayListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHolidayListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHolidayListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHolidayListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHolidayListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHolidayListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHolidayListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHoliday is called when production holiday is entered.
func (s *BaseHolidayListener) EnterHoliday(ctx *HolidayContext) {}

// ExitHoliday is called when production holiday is exited.
func (s *BaseHolidayListener) ExitHoliday(ctx *HolidayContext) {}
