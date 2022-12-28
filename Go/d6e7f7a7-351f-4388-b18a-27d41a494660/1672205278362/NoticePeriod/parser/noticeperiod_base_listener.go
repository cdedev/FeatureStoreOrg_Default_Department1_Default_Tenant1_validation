// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205278362/NoticePeriod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NoticePeriod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNoticePeriodListener is a complete listener for a parse tree produced by NoticePeriodParser.
type BaseNoticePeriodListener struct{}

var _ NoticePeriodListener = &BaseNoticePeriodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNoticePeriodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNoticePeriodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNoticePeriodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNoticePeriodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNoticePeriodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNoticePeriodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNoticeperiod is called when production noticeperiod is entered.
func (s *BaseNoticePeriodListener) EnterNoticeperiod(ctx *NoticeperiodContext) {}

// ExitNoticeperiod is called when production noticeperiod is exited.
func (s *BaseNoticePeriodListener) ExitNoticeperiod(ctx *NoticeperiodContext) {}
