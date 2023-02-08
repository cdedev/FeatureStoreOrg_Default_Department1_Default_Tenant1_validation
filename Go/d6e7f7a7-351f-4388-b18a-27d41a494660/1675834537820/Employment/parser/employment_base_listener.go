// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834537820/Employment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employment

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmploymentListener is a complete listener for a parse tree produced by EmploymentParser.
type BaseEmploymentListener struct{}

var _ EmploymentListener = &BaseEmploymentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmploymentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmploymentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmploymentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmploymentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmploymentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmploymentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmployment is called when production employment is entered.
func (s *BaseEmploymentListener) EnterEmployment(ctx *EmploymentContext) {}

// ExitEmployment is called when production employment is exited.
func (s *BaseEmploymentListener) ExitEmployment(ctx *EmploymentContext) {}
