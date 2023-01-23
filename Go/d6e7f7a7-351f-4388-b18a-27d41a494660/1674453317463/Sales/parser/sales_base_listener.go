// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674453317463/Sales.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sales

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSalesListener is a complete listener for a parse tree produced by SalesParser.
type BaseSalesListener struct{}

var _ SalesListener = &BaseSalesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSalesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSalesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSalesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSalesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSalesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSalesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSales is called when production sales is entered.
func (s *BaseSalesListener) EnterSales(ctx *SalesContext) {}

// ExitSales is called when production sales is exited.
func (s *BaseSalesListener) ExitSales(ctx *SalesContext) {}
