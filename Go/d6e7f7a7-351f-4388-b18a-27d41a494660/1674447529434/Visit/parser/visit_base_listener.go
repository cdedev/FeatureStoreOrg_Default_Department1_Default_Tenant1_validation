// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447529434/Visit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVisitListener is a complete listener for a parse tree produced by VisitParser.
type BaseVisitListener struct{}

var _ VisitListener = &BaseVisitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVisitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVisitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVisitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVisitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVisitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVisitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVisit is called when production visit is entered.
func (s *BaseVisitListener) EnterVisit(ctx *VisitContext) {}

// ExitVisit is called when production visit is exited.
func (s *BaseVisitListener) ExitVisit(ctx *VisitContext) {}
