// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867422033/ExchUsd.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ExchUsd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExchUsdListener is a complete listener for a parse tree produced by ExchUsdParser.
type BaseExchUsdListener struct{}

var _ ExchUsdListener = &BaseExchUsdListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExchUsdListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExchUsdListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExchUsdListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExchUsdListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExchUsdListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExchUsdListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExchusd is called when production exchusd is entered.
func (s *BaseExchUsdListener) EnterExchusd(ctx *ExchusdContext) {}

// ExitExchusd is called when production exchusd is exited.
func (s *BaseExchUsdListener) ExitExchusd(ctx *ExchusdContext) {}
