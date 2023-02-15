// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676462327753/WorkingCapacity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WorkingCapacity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWorkingCapacityListener is a complete listener for a parse tree produced by WorkingCapacityParser.
type BaseWorkingCapacityListener struct{}

var _ WorkingCapacityListener = &BaseWorkingCapacityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWorkingCapacityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWorkingCapacityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWorkingCapacityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWorkingCapacityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWorkingCapacityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWorkingCapacityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWorkingcapacity is called when production workingcapacity is entered.
func (s *BaseWorkingCapacityListener) EnterWorkingcapacity(ctx *WorkingcapacityContext) {}

// ExitWorkingcapacity is called when production workingcapacity is exited.
func (s *BaseWorkingCapacityListener) ExitWorkingcapacity(ctx *WorkingcapacityContext) {}
