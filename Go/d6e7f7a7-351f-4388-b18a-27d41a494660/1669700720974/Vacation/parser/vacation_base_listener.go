// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700720974/Vacation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vacation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVacationListener is a complete listener for a parse tree produced by VacationParser.
type BaseVacationListener struct{}

var _ VacationListener = &BaseVacationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVacationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVacationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVacationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVacationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVacationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVacationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVacation is called when production vacation is entered.
func (s *BaseVacationListener) EnterVacation(ctx *VacationContext) {}

// ExitVacation is called when production vacation is exited.
func (s *BaseVacationListener) ExitVacation(ctx *VacationContext) {}
