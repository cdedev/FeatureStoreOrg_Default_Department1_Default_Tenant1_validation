// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672114544968/Interval.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Interval

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntervalListener is a complete listener for a parse tree produced by IntervalParser.
type BaseIntervalListener struct{}

var _ IntervalListener = &BaseIntervalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntervalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntervalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntervalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntervalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntervalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntervalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseIntervalListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseIntervalListener) ExitInterval(ctx *IntervalContext) {}
