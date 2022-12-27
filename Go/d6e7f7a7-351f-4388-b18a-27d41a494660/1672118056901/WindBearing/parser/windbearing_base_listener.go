// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118056901/WindBearing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WindBearing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWindBearingListener is a complete listener for a parse tree produced by WindBearingParser.
type BaseWindBearingListener struct{}

var _ WindBearingListener = &BaseWindBearingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWindBearingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWindBearingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWindBearingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWindBearingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWindBearingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWindBearingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWindbearing is called when production windbearing is entered.
func (s *BaseWindBearingListener) EnterWindbearing(ctx *WindbearingContext) {}

// ExitWindbearing is called when production windbearing is exited.
func (s *BaseWindBearingListener) ExitWindbearing(ctx *WindbearingContext) {}
