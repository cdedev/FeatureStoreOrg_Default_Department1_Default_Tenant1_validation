// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719178745/Industry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Industry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIndustryListener is a complete listener for a parse tree produced by IndustryParser.
type BaseIndustryListener struct{}

var _ IndustryListener = &BaseIndustryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIndustryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIndustryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIndustryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIndustryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIndustryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIndustryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIndustry is called when production industry is entered.
func (s *BaseIndustryListener) EnterIndustry(ctx *IndustryContext) {}

// ExitIndustry is called when production industry is exited.
func (s *BaseIndustryListener) ExitIndustry(ctx *IndustryContext) {}
