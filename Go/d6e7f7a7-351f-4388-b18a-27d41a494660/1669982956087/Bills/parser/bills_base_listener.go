// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982956087/Bills.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bills

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBillsListener is a complete listener for a parse tree produced by BillsParser.
type BaseBillsListener struct{}

var _ BillsListener = &BaseBillsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBillsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBillsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBillsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBillsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBillsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBillsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBills is called when production bills is entered.
func (s *BaseBillsListener) EnterBills(ctx *BillsContext) {}

// ExitBills is called when production bills is exited.
func (s *BaseBillsListener) ExitBills(ctx *BillsContext) {}
