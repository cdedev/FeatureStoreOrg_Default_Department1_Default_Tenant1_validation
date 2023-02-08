// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845097577/Vibration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vibration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVibrationListener is a complete listener for a parse tree produced by VibrationParser.
type BaseVibrationListener struct{}

var _ VibrationListener = &BaseVibrationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVibrationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVibrationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVibrationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVibrationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVibrationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVibrationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVibration is called when production vibration is entered.
func (s *BaseVibrationListener) EnterVibration(ctx *VibrationContext) {}

// ExitVibration is called when production vibration is exited.
func (s *BaseVibrationListener) ExitVibration(ctx *VibrationContext) {}
