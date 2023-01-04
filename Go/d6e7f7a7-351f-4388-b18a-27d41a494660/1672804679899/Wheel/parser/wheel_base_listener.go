// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804679899/Wheel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wheel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWheelListener is a complete listener for a parse tree produced by WheelParser.
type BaseWheelListener struct{}

var _ WheelListener = &BaseWheelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWheelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWheelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWheelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWheelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWheelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWheelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWheel is called when production wheel is entered.
func (s *BaseWheelListener) EnterWheel(ctx *WheelContext) {}

// ExitWheel is called when production wheel is exited.
func (s *BaseWheelListener) ExitWheel(ctx *WheelContext) {}
