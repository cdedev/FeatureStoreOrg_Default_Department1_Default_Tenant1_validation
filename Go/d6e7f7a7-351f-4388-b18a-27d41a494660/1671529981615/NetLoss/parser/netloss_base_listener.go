// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671529981615/NetLoss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NetLoss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNetLossListener is a complete listener for a parse tree produced by NetLossParser.
type BaseNetLossListener struct{}

var _ NetLossListener = &BaseNetLossListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNetLossListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNetLossListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNetLossListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNetLossListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNetLossListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNetLossListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNetloss is called when production netloss is entered.
func (s *BaseNetLossListener) EnterNetloss(ctx *NetlossContext) {}

// ExitNetloss is called when production netloss is exited.
func (s *BaseNetLossListener) ExitNetloss(ctx *NetlossContext) {}
