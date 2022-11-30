// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780323186/Oil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oil

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOilListener is a complete listener for a parse tree produced by OilParser.
type BaseOilListener struct{}

var _ OilListener = &BaseOilListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOilListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOilListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOilListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOilListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOilListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOilListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOil is called when production oil is entered.
func (s *BaseOilListener) EnterOil(ctx *OilContext) {}

// ExitOil is called when production oil is exited.
func (s *BaseOilListener) ExitOil(ctx *OilContext) {}
