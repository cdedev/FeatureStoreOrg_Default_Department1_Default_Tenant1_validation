// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652889685/BloodPressure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodPressure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBloodPressureListener is a complete listener for a parse tree produced by BloodPressureParser.
type BaseBloodPressureListener struct{}

var _ BloodPressureListener = &BaseBloodPressureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBloodPressureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBloodPressureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBloodPressureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBloodPressureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBloodPressureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBloodPressureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBloodpressure is called when production bloodpressure is entered.
func (s *BaseBloodPressureListener) EnterBloodpressure(ctx *BloodpressureContext) {}

// ExitBloodpressure is called when production bloodpressure is exited.
func (s *BaseBloodPressureListener) ExitBloodpressure(ctx *BloodpressureContext) {}
