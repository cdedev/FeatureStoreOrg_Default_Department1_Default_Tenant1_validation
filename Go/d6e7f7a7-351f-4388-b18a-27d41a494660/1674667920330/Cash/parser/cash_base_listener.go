// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667920330/Cash.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cash

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCashListener is a complete listener for a parse tree produced by CashParser.
type BaseCashListener struct{}

var _ CashListener = &BaseCashListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCashListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCashListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCashListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCashListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCashListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCashListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCash is called when production cash is entered.
func (s *BaseCashListener) EnterCash(ctx *CashContext) {}

// ExitCash is called when production cash is exited.
func (s *BaseCashListener) ExitCash(ctx *CashContext) {}
