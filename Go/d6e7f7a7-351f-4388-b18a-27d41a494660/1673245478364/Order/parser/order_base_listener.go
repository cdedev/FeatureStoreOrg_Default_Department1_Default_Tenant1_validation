// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673245478364/Order.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Order

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOrderListener is a complete listener for a parse tree produced by OrderParser.
type BaseOrderListener struct{}

var _ OrderListener = &BaseOrderListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOrderListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOrderListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOrderListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOrderListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOrderListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOrderListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrder is called when production order is entered.
func (s *BaseOrderListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production order is exited.
func (s *BaseOrderListener) ExitOrder(ctx *OrderContext) {}
