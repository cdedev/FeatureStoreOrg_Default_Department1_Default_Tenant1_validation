// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701718101/Shipping.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shipping

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShippingListener is a complete listener for a parse tree produced by ShippingParser.
type BaseShippingListener struct{}

var _ ShippingListener = &BaseShippingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShippingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShippingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShippingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShippingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseShippingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseShippingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShipping is called when production shipping is entered.
func (s *BaseShippingListener) EnterShipping(ctx *ShippingContext) {}

// ExitShipping is called when production shipping is exited.
func (s *BaseShippingListener) ExitShipping(ctx *ShippingContext) {}
