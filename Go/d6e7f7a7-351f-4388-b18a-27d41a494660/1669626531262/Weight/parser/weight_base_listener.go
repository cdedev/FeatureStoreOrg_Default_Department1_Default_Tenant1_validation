// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626531262/Weight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeightListener is a complete listener for a parse tree produced by WeightParser.
type BaseWeightListener struct{}

var _ WeightListener = &BaseWeightListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeightListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeightListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeightListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeightListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeightListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeightListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeight is called when production weight is entered.
func (s *BaseWeightListener) EnterWeight(ctx *WeightContext) {}

// ExitWeight is called when production weight is exited.
func (s *BaseWeightListener) ExitWeight(ctx *WeightContext) {}
