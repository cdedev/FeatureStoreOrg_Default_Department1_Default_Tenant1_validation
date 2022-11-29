// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700956001/Steps.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Steps

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStepsListener is a complete listener for a parse tree produced by StepsParser.
type BaseStepsListener struct{}

var _ StepsListener = &BaseStepsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStepsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStepsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStepsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStepsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStepsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStepsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSteps is called when production steps is entered.
func (s *BaseStepsListener) EnterSteps(ctx *StepsContext) {}

// ExitSteps is called when production steps is exited.
func (s *BaseStepsListener) ExitSteps(ctx *StepsContext) {}
