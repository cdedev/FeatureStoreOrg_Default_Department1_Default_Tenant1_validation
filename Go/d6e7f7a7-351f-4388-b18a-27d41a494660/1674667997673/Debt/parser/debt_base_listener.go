// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667997673/Debt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDebtListener is a complete listener for a parse tree produced by DebtParser.
type BaseDebtListener struct{}

var _ DebtListener = &BaseDebtListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDebtListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDebtListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDebtListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDebtListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDebtListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDebtListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDebt is called when production debt is entered.
func (s *BaseDebtListener) EnterDebt(ctx *DebtContext) {}

// ExitDebt is called when production debt is exited.
func (s *BaseDebtListener) ExitDebt(ctx *DebtContext) {}
