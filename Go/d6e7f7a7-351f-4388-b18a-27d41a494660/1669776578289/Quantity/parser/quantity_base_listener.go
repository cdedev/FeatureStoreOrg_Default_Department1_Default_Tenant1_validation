// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669776578289/Quantity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quantity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQuantityListener is a complete listener for a parse tree produced by QuantityParser.
type BaseQuantityListener struct{}

var _ QuantityListener = &BaseQuantityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQuantityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQuantityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQuantityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQuantityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseQuantityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseQuantityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BaseQuantityListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BaseQuantityListener) ExitQuantity(ctx *QuantityContext) {}
