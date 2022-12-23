// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769168229/ReportingPeriod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReportingPeriod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReportingPeriodListener is a complete listener for a parse tree produced by ReportingPeriodParser.
type BaseReportingPeriodListener struct{}

var _ ReportingPeriodListener = &BaseReportingPeriodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReportingPeriodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReportingPeriodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReportingPeriodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReportingPeriodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseReportingPeriodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseReportingPeriodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterReportingperiod is called when production reportingperiod is entered.
func (s *BaseReportingPeriodListener) EnterReportingperiod(ctx *ReportingperiodContext) {}

// ExitReportingperiod is called when production reportingperiod is exited.
func (s *BaseReportingPeriodListener) ExitReportingperiod(ctx *ReportingperiodContext) {}
