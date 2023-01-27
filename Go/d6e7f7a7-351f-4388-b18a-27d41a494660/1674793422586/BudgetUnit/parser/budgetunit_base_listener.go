// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793422586/BudgetUnit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BudgetUnit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBudgetUnitListener is a complete listener for a parse tree produced by BudgetUnitParser.
type BaseBudgetUnitListener struct{}

var _ BudgetUnitListener = &BaseBudgetUnitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBudgetUnitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBudgetUnitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBudgetUnitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBudgetUnitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBudgetUnitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBudgetUnitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBudgetunit is called when production budgetunit is entered.
func (s *BaseBudgetUnitListener) EnterBudgetunit(ctx *BudgetunitContext) {}

// ExitBudgetunit is called when production budgetunit is exited.
func (s *BaseBudgetUnitListener) ExitBudgetunit(ctx *BudgetunitContext) {}
