// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672285713110/RevisedAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RevisedAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRevisedAmountListener is a complete listener for a parse tree produced by RevisedAmountParser.
type BaseRevisedAmountListener struct{}

var _ RevisedAmountListener = &BaseRevisedAmountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRevisedAmountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRevisedAmountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRevisedAmountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRevisedAmountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRevisedAmountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRevisedAmountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRevisedamount is called when production revisedamount is entered.
func (s *BaseRevisedAmountListener) EnterRevisedamount(ctx *RevisedamountContext) {}

// ExitRevisedamount is called when production revisedamount is exited.
func (s *BaseRevisedAmountListener) ExitRevisedamount(ctx *RevisedamountContext) {}
