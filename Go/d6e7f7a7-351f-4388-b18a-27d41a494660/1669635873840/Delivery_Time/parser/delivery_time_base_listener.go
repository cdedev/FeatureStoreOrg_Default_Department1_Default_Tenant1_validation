// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669635873840/Delivery_Time.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delivery_Time

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDelivery_TimeListener is a complete listener for a parse tree produced by Delivery_TimeParser.
type BaseDelivery_TimeListener struct{}

var _ Delivery_TimeListener = &BaseDelivery_TimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDelivery_TimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDelivery_TimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDelivery_TimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDelivery_TimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDelivery_TimeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDelivery_TimeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDelivery_time is called when production delivery_time is entered.
func (s *BaseDelivery_TimeListener) EnterDelivery_time(ctx *Delivery_timeContext) {}

// ExitDelivery_time is called when production delivery_time is exited.
func (s *BaseDelivery_TimeListener) ExitDelivery_time(ctx *Delivery_timeContext) {}
