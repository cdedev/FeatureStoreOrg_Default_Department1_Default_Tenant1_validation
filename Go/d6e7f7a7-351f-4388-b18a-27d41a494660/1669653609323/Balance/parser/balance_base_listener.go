// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653609323/Balance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Balance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBalanceListener is a complete listener for a parse tree produced by BalanceParser.
type BaseBalanceListener struct{}

var _ BalanceListener = &BaseBalanceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBalanceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBalanceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBalanceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBalanceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBalanceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBalanceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBalance is called when production balance is entered.
func (s *BaseBalanceListener) EnterBalance(ctx *BalanceContext) {}

// ExitBalance is called when production balance is exited.
func (s *BaseBalanceListener) ExitBalance(ctx *BalanceContext) {}
