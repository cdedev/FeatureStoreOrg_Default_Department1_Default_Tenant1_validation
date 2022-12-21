// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671602238737/Loan.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLoanListener is a complete listener for a parse tree produced by LoanParser.
type BaseLoanListener struct{}

var _ LoanListener = &BaseLoanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLoanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLoanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLoanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLoanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLoanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLoanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLoan is called when production loan is entered.
func (s *BaseLoanListener) EnterLoan(ctx *LoanContext) {}

// ExitLoan is called when production loan is exited.
func (s *BaseLoanListener) ExitLoan(ctx *LoanContext) {}
