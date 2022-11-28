// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669630209027/Loss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLossListener is a complete listener for a parse tree produced by LossParser.
type BaseLossListener struct{}

var _ LossListener = &BaseLossListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLossListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLossListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLossListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLossListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLossListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLossListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLoss is called when production loss is entered.
func (s *BaseLossListener) EnterLoss(ctx *LossContext) {}

// ExitLoss is called when production loss is exited.
func (s *BaseLossListener) ExitLoss(ctx *LossContext) {}
