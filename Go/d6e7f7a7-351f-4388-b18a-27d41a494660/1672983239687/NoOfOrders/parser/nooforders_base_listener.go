// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983239687/NoOfOrders.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NoOfOrders

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNoOfOrdersListener is a complete listener for a parse tree produced by NoOfOrdersParser.
type BaseNoOfOrdersListener struct{}

var _ NoOfOrdersListener = &BaseNoOfOrdersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNoOfOrdersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNoOfOrdersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNoOfOrdersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNoOfOrdersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNoOfOrdersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNoOfOrdersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNooforders is called when production nooforders is entered.
func (s *BaseNoOfOrdersListener) EnterNooforders(ctx *NoofordersContext) {}

// ExitNooforders is called when production nooforders is exited.
func (s *BaseNoOfOrdersListener) ExitNooforders(ctx *NoofordersContext) {}
