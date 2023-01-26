// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719222532/Mortgage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortgage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMortgageListener is a complete listener for a parse tree produced by MortgageParser.
type BaseMortgageListener struct{}

var _ MortgageListener = &BaseMortgageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMortgageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMortgageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMortgageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMortgageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMortgageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMortgageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMortgage is called when production mortgage is entered.
func (s *BaseMortgageListener) EnterMortgage(ctx *MortgageContext) {}

// ExitMortgage is called when production mortgage is exited.
func (s *BaseMortgageListener) ExitMortgage(ctx *MortgageContext) {}
