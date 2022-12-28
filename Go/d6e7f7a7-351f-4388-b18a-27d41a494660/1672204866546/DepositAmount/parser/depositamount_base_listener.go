// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672204866546/DepositAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DepositAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDepositAmountListener is a complete listener for a parse tree produced by DepositAmountParser.
type BaseDepositAmountListener struct{}

var _ DepositAmountListener = &BaseDepositAmountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDepositAmountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDepositAmountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDepositAmountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDepositAmountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDepositAmountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDepositAmountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDepositamount is called when production depositamount is entered.
func (s *BaseDepositAmountListener) EnterDepositamount(ctx *DepositamountContext) {}

// ExitDepositamount is called when production depositamount is exited.
func (s *BaseDepositAmountListener) ExitDepositamount(ctx *DepositamountContext) {}
