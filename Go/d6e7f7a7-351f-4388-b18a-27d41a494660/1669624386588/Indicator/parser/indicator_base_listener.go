// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624386588/Indicator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Indicator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIndicatorListener is a complete listener for a parse tree produced by IndicatorParser.
type BaseIndicatorListener struct{}

var _ IndicatorListener = &BaseIndicatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIndicatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIndicatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIndicatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIndicatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIndicatorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIndicatorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIndicator is called when production indicator is entered.
func (s *BaseIndicatorListener) EnterIndicator(ctx *IndicatorContext) {}

// ExitIndicator is called when production indicator is exited.
func (s *BaseIndicatorListener) ExitIndicator(ctx *IndicatorContext) {}
