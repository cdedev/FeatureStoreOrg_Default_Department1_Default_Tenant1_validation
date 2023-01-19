// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096918596/Roll.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Roll

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRollListener is a complete listener for a parse tree produced by RollParser.
type BaseRollListener struct{}

var _ RollListener = &BaseRollListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRollListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRollListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRollListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRollListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRollListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRollListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRoll is called when production roll is entered.
func (s *BaseRollListener) EnterRoll(ctx *RollContext) {}

// ExitRoll is called when production roll is exited.
func (s *BaseRollListener) ExitRoll(ctx *RollContext) {}
