// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655797520/Heartrate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Heartrate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHeartrateListener is a complete listener for a parse tree produced by HeartrateParser.
type BaseHeartrateListener struct{}

var _ HeartrateListener = &BaseHeartrateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHeartrateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHeartrateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHeartrateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHeartrateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHeartrateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHeartrateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHeartrate is called when production heartrate is entered.
func (s *BaseHeartrateListener) EnterHeartrate(ctx *HeartrateContext) {}

// ExitHeartrate is called when production heartrate is exited.
func (s *BaseHeartrateListener) ExitHeartrate(ctx *HeartrateContext) {}
