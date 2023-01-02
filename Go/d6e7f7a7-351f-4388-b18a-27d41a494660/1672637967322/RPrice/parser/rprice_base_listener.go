// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637967322/RPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RPrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRPriceListener is a complete listener for a parse tree produced by RPriceParser.
type BaseRPriceListener struct{}

var _ RPriceListener = &BaseRPriceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRPriceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRPriceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRPriceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRPriceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRPriceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRPriceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRprice is called when production rprice is entered.
func (s *BaseRPriceListener) EnterRprice(ctx *RpriceContext) {}

// ExitRprice is called when production rprice is exited.
func (s *BaseRPriceListener) ExitRprice(ctx *RpriceContext) {}
