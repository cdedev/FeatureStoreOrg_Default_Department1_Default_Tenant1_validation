// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884259507/Revenue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Revenue

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRevenueListener is a complete listener for a parse tree produced by RevenueParser.
type BaseRevenueListener struct{}

var _ RevenueListener = &BaseRevenueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRevenueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRevenueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRevenueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRevenueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRevenueListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRevenueListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRevenue is called when production revenue is entered.
func (s *BaseRevenueListener) EnterRevenue(ctx *RevenueContext) {}

// ExitRevenue is called when production revenue is exited.
func (s *BaseRevenueListener) ExitRevenue(ctx *RevenueContext) {}
