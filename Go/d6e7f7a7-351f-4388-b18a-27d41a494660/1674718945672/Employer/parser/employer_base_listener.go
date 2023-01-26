// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674718945672/Employer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmployerListener is a complete listener for a parse tree produced by EmployerParser.
type BaseEmployerListener struct{}

var _ EmployerListener = &BaseEmployerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmployerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmployerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmployerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmployerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmployerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmployerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmployer is called when production employer is entered.
func (s *BaseEmployerListener) EnterEmployer(ctx *EmployerContext) {}

// ExitEmployer is called when production employer is exited.
func (s *BaseEmployerListener) ExitEmployer(ctx *EmployerContext) {}
