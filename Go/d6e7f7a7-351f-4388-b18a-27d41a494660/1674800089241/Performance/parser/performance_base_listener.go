// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674800089241/Performance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Performance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePerformanceListener is a complete listener for a parse tree produced by PerformanceParser.
type BasePerformanceListener struct{}

var _ PerformanceListener = &BasePerformanceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePerformanceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePerformanceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePerformanceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePerformanceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePerformanceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePerformanceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPerformance is called when production performance is entered.
func (s *BasePerformanceListener) EnterPerformance(ctx *PerformanceContext) {}

// ExitPerformance is called when production performance is exited.
func (s *BasePerformanceListener) ExitPerformance(ctx *PerformanceContext) {}
