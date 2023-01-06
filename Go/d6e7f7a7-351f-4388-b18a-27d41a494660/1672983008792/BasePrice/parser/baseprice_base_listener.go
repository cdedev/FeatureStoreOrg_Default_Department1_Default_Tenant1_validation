// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983008792/BasePrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BasePrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBasePriceListener is a complete listener for a parse tree produced by BasePriceParser.
type BaseBasePriceListener struct{}

var _ BasePriceListener = &BaseBasePriceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBasePriceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBasePriceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBasePriceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBasePriceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBasePriceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBasePriceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBaseprice is called when production baseprice is entered.
func (s *BaseBasePriceListener) EnterBaseprice(ctx *BasepriceContext) {}

// ExitBaseprice is called when production baseprice is exited.
func (s *BaseBasePriceListener) ExitBaseprice(ctx *BasepriceContext) {}
