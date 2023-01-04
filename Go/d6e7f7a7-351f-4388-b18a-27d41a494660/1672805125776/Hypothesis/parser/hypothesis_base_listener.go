// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672805125776/Hypothesis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hypothesis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHypothesisListener is a complete listener for a parse tree produced by HypothesisParser.
type BaseHypothesisListener struct{}

var _ HypothesisListener = &BaseHypothesisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHypothesisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHypothesisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHypothesisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHypothesisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHypothesisListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHypothesisListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHypothesis is called when production hypothesis is entered.
func (s *BaseHypothesisListener) EnterHypothesis(ctx *HypothesisContext) {}

// ExitHypothesis is called when production hypothesis is exited.
func (s *BaseHypothesisListener) ExitHypothesis(ctx *HypothesisContext) {}
