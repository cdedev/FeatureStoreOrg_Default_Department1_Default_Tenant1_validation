// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675220837600/Payment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Payment

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePaymentListener is a complete listener for a parse tree produced by PaymentParser.
type BasePaymentListener struct{}

var _ PaymentListener = &BasePaymentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePaymentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePaymentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePaymentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePaymentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePaymentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePaymentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPayment is called when production payment is entered.
func (s *BasePaymentListener) EnterPayment(ctx *PaymentContext) {}

// ExitPayment is called when production payment is exited.
func (s *BasePaymentListener) ExitPayment(ctx *PaymentContext) {}
