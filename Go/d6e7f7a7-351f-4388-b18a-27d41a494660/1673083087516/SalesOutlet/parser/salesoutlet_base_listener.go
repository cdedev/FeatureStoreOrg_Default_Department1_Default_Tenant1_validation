// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083087516/SalesOutlet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SalesOutlet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSalesOutletListener is a complete listener for a parse tree produced by SalesOutletParser.
type BaseSalesOutletListener struct{}

var _ SalesOutletListener = &BaseSalesOutletListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSalesOutletListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSalesOutletListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSalesOutletListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSalesOutletListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSalesOutletListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSalesOutletListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSalesoutlet is called when production salesoutlet is entered.
func (s *BaseSalesOutletListener) EnterSalesoutlet(ctx *SalesoutletContext) {}

// ExitSalesoutlet is called when production salesoutlet is exited.
func (s *BaseSalesOutletListener) ExitSalesoutlet(ctx *SalesoutletContext) {}
