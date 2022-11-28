// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624662791/Time.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTimeListener is a complete listener for a parse tree produced by TimeParser.
type BaseTimeListener struct{}

var _ TimeListener = &BaseTimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTimeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTimeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTime is called when production time is entered.
func (s *BaseTimeListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseTimeListener) ExitTime(ctx *TimeContext) {}
