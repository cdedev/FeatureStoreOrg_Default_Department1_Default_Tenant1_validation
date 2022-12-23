// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769755680/GrossAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GrossAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrossAmountListener is a complete listener for a parse tree produced by GrossAmountParser.
type BaseGrossAmountListener struct{}

var _ GrossAmountListener = &BaseGrossAmountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrossAmountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrossAmountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrossAmountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrossAmountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGrossAmountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGrossAmountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGrossamount is called when production grossamount is entered.
func (s *BaseGrossAmountListener) EnterGrossamount(ctx *GrossamountContext) {}

// ExitGrossamount is called when production grossamount is exited.
func (s *BaseGrossAmountListener) ExitGrossamount(ctx *GrossamountContext) {}
