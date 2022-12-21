// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671600641537/MarketPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarketPrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarketPriceListener is a complete listener for a parse tree produced by MarketPriceParser.
type BaseMarketPriceListener struct{}

var _ MarketPriceListener = &BaseMarketPriceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarketPriceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarketPriceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarketPriceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarketPriceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarketPriceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarketPriceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarketprice is called when production marketprice is entered.
func (s *BaseMarketPriceListener) EnterMarketprice(ctx *MarketpriceContext) {}

// ExitMarketprice is called when production marketprice is exited.
func (s *BaseMarketPriceListener) ExitMarketprice(ctx *MarketpriceContext) {}
