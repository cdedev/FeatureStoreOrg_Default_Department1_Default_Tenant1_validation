// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669973382517/Inflation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Inflation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInflationListener is a complete listener for a parse tree produced by InflationParser.
type BaseInflationListener struct{}

var _ InflationListener = &BaseInflationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInflationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInflationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInflationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInflationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInflationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInflationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInflation is called when production inflation is entered.
func (s *BaseInflationListener) EnterInflation(ctx *InflationContext) {}

// ExitInflation is called when production inflation is exited.
func (s *BaseInflationListener) ExitInflation(ctx *InflationContext) {}
