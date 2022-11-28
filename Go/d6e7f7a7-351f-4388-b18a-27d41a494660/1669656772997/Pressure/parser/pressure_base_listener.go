// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656772997/Pressure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pressure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePressureListener is a complete listener for a parse tree produced by PressureParser.
type BasePressureListener struct{}

var _ PressureListener = &BasePressureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePressureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePressureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePressureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePressureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePressureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePressureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPressure is called when production pressure is entered.
func (s *BasePressureListener) EnterPressure(ctx *PressureContext) {}

// ExitPressure is called when production pressure is exited.
func (s *BasePressureListener) ExitPressure(ctx *PressureContext) {}
