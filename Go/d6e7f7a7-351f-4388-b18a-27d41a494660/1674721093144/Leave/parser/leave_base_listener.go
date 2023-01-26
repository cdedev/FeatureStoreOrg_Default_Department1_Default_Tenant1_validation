// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721093144/Leave.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Leave

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLeaveListener is a complete listener for a parse tree produced by LeaveParser.
type BaseLeaveListener struct{}

var _ LeaveListener = &BaseLeaveListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLeaveListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLeaveListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLeaveListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLeaveListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLeaveListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLeaveListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLeave is called when production leave is entered.
func (s *BaseLeaveListener) EnterLeave(ctx *LeaveContext) {}

// ExitLeave is called when production leave is exited.
func (s *BaseLeaveListener) ExitLeave(ctx *LeaveContext) {}
