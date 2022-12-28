// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672198512597/OrderTotal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OrderTotal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOrderTotalListener is a complete listener for a parse tree produced by OrderTotalParser.
type BaseOrderTotalListener struct{}

var _ OrderTotalListener = &BaseOrderTotalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOrderTotalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOrderTotalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOrderTotalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOrderTotalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOrderTotalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOrderTotalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrdertotal is called when production ordertotal is entered.
func (s *BaseOrderTotalListener) EnterOrdertotal(ctx *OrdertotalContext) {}

// ExitOrdertotal is called when production ordertotal is exited.
func (s *BaseOrderTotalListener) ExitOrdertotal(ctx *OrdertotalContext) {}
