// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377471694/ClockSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ClockSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClockSpeedListener is a complete listener for a parse tree produced by ClockSpeedParser.
type BaseClockSpeedListener struct{}

var _ ClockSpeedListener = &BaseClockSpeedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClockSpeedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClockSpeedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClockSpeedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClockSpeedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseClockSpeedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseClockSpeedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterClockspeed is called when production clockspeed is entered.
func (s *BaseClockSpeedListener) EnterClockspeed(ctx *ClockspeedContext) {}

// ExitClockspeed is called when production clockspeed is exited.
func (s *BaseClockSpeedListener) ExitClockspeed(ctx *ClockspeedContext) {}
