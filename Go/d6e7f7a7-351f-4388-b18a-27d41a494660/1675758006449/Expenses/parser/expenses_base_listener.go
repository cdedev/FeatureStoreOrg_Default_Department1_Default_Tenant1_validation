// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758006449/Expenses.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenses

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpensesListener is a complete listener for a parse tree produced by ExpensesParser.
type BaseExpensesListener struct{}

var _ ExpensesListener = &BaseExpensesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpensesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpensesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpensesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpensesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpensesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpensesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpenses is called when production expenses is entered.
func (s *BaseExpensesListener) EnterExpenses(ctx *ExpensesContext) {}

// ExitExpenses is called when production expenses is exited.
func (s *BaseExpensesListener) ExitExpenses(ctx *ExpensesContext) {}
