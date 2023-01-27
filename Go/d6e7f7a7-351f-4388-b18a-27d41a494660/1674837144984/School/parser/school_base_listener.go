// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837144984/School.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // School

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSchoolListener is a complete listener for a parse tree produced by SchoolParser.
type BaseSchoolListener struct{}

var _ SchoolListener = &BaseSchoolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSchoolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSchoolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSchoolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSchoolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSchoolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSchoolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSchool is called when production school is entered.
func (s *BaseSchoolListener) EnterSchool(ctx *SchoolContext) {}

// ExitSchool is called when production school is exited.
func (s *BaseSchoolListener) ExitSchool(ctx *SchoolContext) {}
