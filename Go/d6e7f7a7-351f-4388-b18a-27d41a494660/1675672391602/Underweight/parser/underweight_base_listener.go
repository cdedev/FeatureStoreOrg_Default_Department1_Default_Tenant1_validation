// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675672391602/Underweight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Underweight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUnderweightListener is a complete listener for a parse tree produced by UnderweightParser.
type BaseUnderweightListener struct{}

var _ UnderweightListener = &BaseUnderweightListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUnderweightListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUnderweightListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUnderweightListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUnderweightListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUnderweightListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUnderweightListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnderweight is called when production underweight is entered.
func (s *BaseUnderweightListener) EnterUnderweight(ctx *UnderweightContext) {}

// ExitUnderweight is called when production underweight is exited.
func (s *BaseUnderweightListener) ExitUnderweight(ctx *UnderweightContext) {}
