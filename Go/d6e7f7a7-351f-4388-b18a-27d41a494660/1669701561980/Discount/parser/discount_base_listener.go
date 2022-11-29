// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701561980/Discount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Discount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiscountListener is a complete listener for a parse tree produced by DiscountParser.
type BaseDiscountListener struct{}

var _ DiscountListener = &BaseDiscountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiscountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiscountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiscountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiscountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiscountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiscountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDiscount is called when production discount is entered.
func (s *BaseDiscountListener) EnterDiscount(ctx *DiscountContext) {}

// ExitDiscount is called when production discount is exited.
func (s *BaseDiscountListener) ExitDiscount(ctx *DiscountContext) {}
