// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673250498849/Subject.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Subject

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSubjectListener is a complete listener for a parse tree produced by SubjectParser.
type BaseSubjectListener struct{}

var _ SubjectListener = &BaseSubjectListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSubjectListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSubjectListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSubjectListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSubjectListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSubjectListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSubjectListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubject is called when production subject is entered.
func (s *BaseSubjectListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BaseSubjectListener) ExitSubject(ctx *SubjectContext) {}
