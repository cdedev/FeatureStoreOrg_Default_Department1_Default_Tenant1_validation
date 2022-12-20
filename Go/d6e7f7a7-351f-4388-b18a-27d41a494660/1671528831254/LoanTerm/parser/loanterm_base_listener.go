// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528831254/LoanTerm.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LoanTerm

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLoanTermListener is a complete listener for a parse tree produced by LoanTermParser.
type BaseLoanTermListener struct{}

var _ LoanTermListener = &BaseLoanTermListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLoanTermListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLoanTermListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLoanTermListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLoanTermListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLoanTermListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLoanTermListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLoanterm is called when production loanterm is entered.
func (s *BaseLoanTermListener) EnterLoanterm(ctx *LoantermContext) {}

// ExitLoanterm is called when production loanterm is exited.
func (s *BaseLoanTermListener) ExitLoanterm(ctx *LoantermContext) {}
