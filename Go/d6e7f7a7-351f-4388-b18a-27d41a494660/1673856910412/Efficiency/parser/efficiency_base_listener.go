// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673856910412/Efficiency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Efficiency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEfficiencyListener is a complete listener for a parse tree produced by EfficiencyParser.
type BaseEfficiencyListener struct{}

var _ EfficiencyListener = &BaseEfficiencyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEfficiencyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEfficiencyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEfficiencyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEfficiencyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEfficiencyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEfficiencyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEfficiency is called when production efficiency is entered.
func (s *BaseEfficiencyListener) EnterEfficiency(ctx *EfficiencyContext) {}

// ExitEfficiency is called when production efficiency is exited.
func (s *BaseEfficiencyListener) ExitEfficiency(ctx *EfficiencyContext) {}
