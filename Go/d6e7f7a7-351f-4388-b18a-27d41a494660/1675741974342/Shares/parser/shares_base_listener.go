// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675741974342/Shares.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shares

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSharesListener is a complete listener for a parse tree produced by SharesParser.
type BaseSharesListener struct{}

var _ SharesListener = &BaseSharesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSharesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSharesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSharesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSharesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSharesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSharesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShares is called when production shares is entered.
func (s *BaseSharesListener) EnterShares(ctx *SharesContext) {}

// ExitShares is called when production shares is exited.
func (s *BaseSharesListener) ExitShares(ctx *SharesContext) {}
