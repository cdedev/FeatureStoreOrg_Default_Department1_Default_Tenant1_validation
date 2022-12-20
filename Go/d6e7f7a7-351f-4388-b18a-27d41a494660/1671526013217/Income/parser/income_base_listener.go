// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671526013217/Income.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Income

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIncomeListener is a complete listener for a parse tree produced by IncomeParser.
type BaseIncomeListener struct{}

var _ IncomeListener = &BaseIncomeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIncomeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIncomeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIncomeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIncomeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIncomeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIncomeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIncome is called when production income is entered.
func (s *BaseIncomeListener) EnterIncome(ctx *IncomeContext) {}

// ExitIncome is called when production income is exited.
func (s *BaseIncomeListener) ExitIncome(ctx *IncomeContext) {}
