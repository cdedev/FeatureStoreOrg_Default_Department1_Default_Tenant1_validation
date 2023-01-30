// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069318693/Course.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Course

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCourseListener is a complete listener for a parse tree produced by CourseParser.
type BaseCourseListener struct{}

var _ CourseListener = &BaseCourseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCourseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCourseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCourseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCourseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCourseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCourseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCourse is called when production course is entered.
func (s *BaseCourseListener) EnterCourse(ctx *CourseContext) {}

// ExitCourse is called when production course is exited.
func (s *BaseCourseListener) ExitCourse(ctx *CourseContext) {}
