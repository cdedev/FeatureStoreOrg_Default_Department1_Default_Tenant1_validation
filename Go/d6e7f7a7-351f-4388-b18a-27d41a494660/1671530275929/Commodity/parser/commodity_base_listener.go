// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671530275929/Commodity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Commodity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCommodityListener is a complete listener for a parse tree produced by CommodityParser.
type BaseCommodityListener struct{}

var _ CommodityListener = &BaseCommodityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCommodityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCommodityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCommodityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCommodityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCommodityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCommodityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCommodity is called when production commodity is entered.
func (s *BaseCommodityListener) EnterCommodity(ctx *CommodityContext) {}

// ExitCommodity is called when production commodity is exited.
func (s *BaseCommodityListener) ExitCommodity(ctx *CommodityContext) {}
