// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672982473127/OrderHourOfDay.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OrderHourOfDay

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOrderHourOfDayListener is a complete listener for a parse tree produced by OrderHourOfDayParser.
type BaseOrderHourOfDayListener struct{}

var _ OrderHourOfDayListener = &BaseOrderHourOfDayListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOrderHourOfDayListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOrderHourOfDayListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOrderHourOfDayListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOrderHourOfDayListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOrderHourOfDayListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOrderHourOfDayListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrderhourofday is called when production orderhourofday is entered.
func (s *BaseOrderHourOfDayListener) EnterOrderhourofday(ctx *OrderhourofdayContext) {}

// ExitOrderhourofday is called when production orderhourofday is exited.
func (s *BaseOrderHourOfDayListener) ExitOrderhourofday(ctx *OrderhourofdayContext) {}
