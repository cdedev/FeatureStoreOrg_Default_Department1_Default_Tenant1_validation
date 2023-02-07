// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675741900018/Employee.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employee

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmployeeListener is a complete listener for a parse tree produced by EmployeeParser.
type BaseEmployeeListener struct{}

var _ EmployeeListener = &BaseEmployeeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmployeeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmployeeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmployeeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmployeeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmployeeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmployeeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmployee is called when production employee is entered.
func (s *BaseEmployeeListener) EnterEmployee(ctx *EmployeeContext) {}

// ExitEmployee is called when production employee is exited.
func (s *BaseEmployeeListener) ExitEmployee(ctx *EmployeeContext) {}
