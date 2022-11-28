// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659215020/Grade.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Grade

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGradeListener is a complete listener for a parse tree produced by GradeParser.
type BaseGradeListener struct{}

var _ GradeListener = &BaseGradeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGradeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGradeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGradeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGradeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGradeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGradeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGrade is called when production grade is entered.
func (s *BaseGradeListener) EnterGrade(ctx *GradeContext) {}

// ExitGrade is called when production grade is exited.
func (s *BaseGradeListener) ExitGrade(ctx *GradeContext) {}
