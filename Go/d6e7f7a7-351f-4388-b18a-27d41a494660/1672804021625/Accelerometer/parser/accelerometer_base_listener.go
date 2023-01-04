// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804021625/Accelerometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Accelerometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAccelerometerListener is a complete listener for a parse tree produced by AccelerometerParser.
type BaseAccelerometerListener struct{}

var _ AccelerometerListener = &BaseAccelerometerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAccelerometerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAccelerometerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAccelerometerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAccelerometerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAccelerometerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAccelerometerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAccelerometer is called when production accelerometer is entered.
func (s *BaseAccelerometerListener) EnterAccelerometer(ctx *AccelerometerContext) {}

// ExitAccelerometer is called when production accelerometer is exited.
func (s *BaseAccelerometerListener) ExitAccelerometer(ctx *AccelerometerContext) {}
