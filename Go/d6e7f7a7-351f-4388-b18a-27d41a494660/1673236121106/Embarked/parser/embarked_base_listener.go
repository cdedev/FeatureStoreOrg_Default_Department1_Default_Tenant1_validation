// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236121106/Embarked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Embarked

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmbarkedListener is a complete listener for a parse tree produced by EmbarkedParser.
type BaseEmbarkedListener struct{}

var _ EmbarkedListener = &BaseEmbarkedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmbarkedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmbarkedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmbarkedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmbarkedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmbarkedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmbarkedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmbarked is called when production embarked is entered.
func (s *BaseEmbarkedListener) EnterEmbarked(ctx *EmbarkedContext) {}

// ExitEmbarked is called when production embarked is exited.
func (s *BaseEmbarkedListener) ExitEmbarked(ctx *EmbarkedContext) {}
