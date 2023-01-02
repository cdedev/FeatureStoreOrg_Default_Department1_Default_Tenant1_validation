// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637859622/RCoil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RCoil

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRCoilListener is a complete listener for a parse tree produced by RCoilParser.
type BaseRCoilListener struct{}

var _ RCoilListener = &BaseRCoilListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRCoilListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRCoilListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRCoilListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRCoilListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRCoilListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRCoilListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRcoil is called when production rcoil is entered.
func (s *BaseRCoilListener) EnterRcoil(ctx *RcoilContext) {}

// ExitRcoil is called when production rcoil is exited.
func (s *BaseRCoilListener) ExitRcoil(ctx *RcoilContext) {}
