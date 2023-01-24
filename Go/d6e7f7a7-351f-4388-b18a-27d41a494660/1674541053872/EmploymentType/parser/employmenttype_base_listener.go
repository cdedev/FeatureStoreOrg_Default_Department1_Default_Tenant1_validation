// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674541053872/EmploymentType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EmploymentType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmploymentTypeListener is a complete listener for a parse tree produced by EmploymentTypeParser.
type BaseEmploymentTypeListener struct{}

var _ EmploymentTypeListener = &BaseEmploymentTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmploymentTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmploymentTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmploymentTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmploymentTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmploymentTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmploymentTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmploymenttype is called when production employmenttype is entered.
func (s *BaseEmploymentTypeListener) EnterEmploymenttype(ctx *EmploymenttypeContext) {}

// ExitEmploymenttype is called when production employmenttype is exited.
func (s *BaseEmploymentTypeListener) ExitEmploymenttype(ctx *EmploymenttypeContext) {}
