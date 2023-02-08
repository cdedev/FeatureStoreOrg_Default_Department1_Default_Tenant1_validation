// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836456873/Period.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Period

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePeriodListener is a complete listener for a parse tree produced by PeriodParser.
type BasePeriodListener struct{}

var _ PeriodListener = &BasePeriodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePeriodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePeriodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePeriodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePeriodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePeriodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePeriodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPeriod is called when production period is entered.
func (s *BasePeriodListener) EnterPeriod(ctx *PeriodContext) {}

// ExitPeriod is called when production period is exited.
func (s *BasePeriodListener) ExitPeriod(ctx *PeriodContext) {}
