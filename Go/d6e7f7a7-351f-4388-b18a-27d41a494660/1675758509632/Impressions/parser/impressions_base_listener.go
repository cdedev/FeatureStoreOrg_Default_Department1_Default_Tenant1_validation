// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758509632/Impressions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Impressions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseImpressionsListener is a complete listener for a parse tree produced by ImpressionsParser.
type BaseImpressionsListener struct{}

var _ ImpressionsListener = &BaseImpressionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseImpressionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseImpressionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseImpressionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseImpressionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseImpressionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseImpressionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterImpressions is called when production impressions is entered.
func (s *BaseImpressionsListener) EnterImpressions(ctx *ImpressionsContext) {}

// ExitImpressions is called when production impressions is exited.
func (s *BaseImpressionsListener) ExitImpressions(ctx *ImpressionsContext) {}
