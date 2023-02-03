// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675404810293/Angle.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Angle

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAngleListener is a complete listener for a parse tree produced by AngleParser.
type BaseAngleListener struct{}

var _ AngleListener = &BaseAngleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAngleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAngleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAngleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAngleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAngleListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAngleListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAngle is called when production angle is entered.
func (s *BaseAngleListener) EnterAngle(ctx *AngleContext) {}

// ExitAngle is called when production angle is exited.
func (s *BaseAngleListener) ExitAngle(ctx *AngleContext) {}
