// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675227175577/FiscalYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FiscalYear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFiscalYearListener is a complete listener for a parse tree produced by FiscalYearParser.
type BaseFiscalYearListener struct{}

var _ FiscalYearListener = &BaseFiscalYearListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFiscalYearListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFiscalYearListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFiscalYearListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFiscalYearListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFiscalYearListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFiscalYearListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFiscalyear is called when production fiscalyear is entered.
func (s *BaseFiscalYearListener) EnterFiscalyear(ctx *FiscalyearContext) {}

// ExitFiscalyear is called when production fiscalyear is exited.
func (s *BaseFiscalYearListener) ExitFiscalyear(ctx *FiscalyearContext) {}
