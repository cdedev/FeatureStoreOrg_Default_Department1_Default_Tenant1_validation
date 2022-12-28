// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202190388/Intelligence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Intelligence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntelligenceListener is a complete listener for a parse tree produced by IntelligenceParser.
type BaseIntelligenceListener struct{}

var _ IntelligenceListener = &BaseIntelligenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntelligenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntelligenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntelligenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntelligenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntelligenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntelligenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIntelligence is called when production intelligence is entered.
func (s *BaseIntelligenceListener) EnterIntelligence(ctx *IntelligenceContext) {}

// ExitIntelligence is called when production intelligence is exited.
func (s *BaseIntelligenceListener) ExitIntelligence(ctx *IntelligenceContext) {}
