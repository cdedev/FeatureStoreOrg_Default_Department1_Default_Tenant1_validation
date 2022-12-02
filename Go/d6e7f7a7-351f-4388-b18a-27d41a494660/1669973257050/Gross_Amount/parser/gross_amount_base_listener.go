// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669973257050/Gross_Amount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gross_Amount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGross_AmountListener is a complete listener for a parse tree produced by Gross_AmountParser.
type BaseGross_AmountListener struct{}

var _ Gross_AmountListener = &BaseGross_AmountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGross_AmountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGross_AmountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGross_AmountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGross_AmountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGross_AmountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGross_AmountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGross_amount is called when production gross_amount is entered.
func (s *BaseGross_AmountListener) EnterGross_amount(ctx *Gross_amountContext) {}

// ExitGross_amount is called when production gross_amount is exited.
func (s *BaseGross_AmountListener) ExitGross_amount(ctx *Gross_amountContext) {}
