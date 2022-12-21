// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601979388/StopDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StopDuration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStopDurationListener is a complete listener for a parse tree produced by StopDurationParser.
type BaseStopDurationListener struct{}

var _ StopDurationListener = &BaseStopDurationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStopDurationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStopDurationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStopDurationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStopDurationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStopDurationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStopDurationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStopduration is called when production stopduration is entered.
func (s *BaseStopDurationListener) EnterStopduration(ctx *StopdurationContext) {}

// ExitStopduration is called when production stopduration is exited.
func (s *BaseStopDurationListener) ExitStopduration(ctx *StopdurationContext) {}
