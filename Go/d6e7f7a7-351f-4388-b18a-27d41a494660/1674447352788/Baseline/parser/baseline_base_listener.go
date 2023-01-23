// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447352788/Baseline.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Baseline

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBaselineListener is a complete listener for a parse tree produced by BaselineParser.
type BaseBaselineListener struct{}

var _ BaselineListener = &BaseBaselineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBaselineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBaselineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBaselineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBaselineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBaselineListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBaselineListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBaseline is called when production baseline is entered.
func (s *BaseBaselineListener) EnterBaseline(ctx *BaselineContext) {}

// ExitBaseline is called when production baseline is exited.
func (s *BaseBaselineListener) ExitBaseline(ctx *BaselineContext) {}
