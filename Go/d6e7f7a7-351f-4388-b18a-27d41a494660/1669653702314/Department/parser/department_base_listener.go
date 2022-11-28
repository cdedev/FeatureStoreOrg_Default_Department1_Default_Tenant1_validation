// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653702314/Department.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Department

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDepartmentListener is a complete listener for a parse tree produced by DepartmentParser.
type BaseDepartmentListener struct{}

var _ DepartmentListener = &BaseDepartmentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDepartmentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDepartmentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDepartmentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDepartmentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDepartmentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDepartmentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDepartment is called when production department is entered.
func (s *BaseDepartmentListener) EnterDepartment(ctx *DepartmentContext) {}

// ExitDepartment is called when production department is exited.
func (s *BaseDepartmentListener) ExitDepartment(ctx *DepartmentContext) {}
