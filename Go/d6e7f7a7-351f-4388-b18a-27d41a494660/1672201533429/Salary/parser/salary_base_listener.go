// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201533429/Salary.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Salary

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSalaryListener is a complete listener for a parse tree produced by SalaryParser.
type BaseSalaryListener struct{}

var _ SalaryListener = &BaseSalaryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSalaryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSalaryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSalaryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSalaryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSalaryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSalaryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSalary is called when production salary is entered.
func (s *BaseSalaryListener) EnterSalary(ctx *SalaryContext) {}

// ExitSalary is called when production salary is exited.
func (s *BaseSalaryListener) ExitSalary(ctx *SalaryContext) {}
