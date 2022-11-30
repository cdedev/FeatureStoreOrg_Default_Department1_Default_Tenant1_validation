// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669786719095/Market.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Market

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarketListener is a complete listener for a parse tree produced by MarketParser.
type BaseMarketListener struct{}

var _ MarketListener = &BaseMarketListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarketListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarketListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarketListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarketListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarketListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarketListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarket is called when production market is entered.
func (s *BaseMarketListener) EnterMarket(ctx *MarketContext) {}

// ExitMarket is called when production market is exited.
func (s *BaseMarketListener) ExitMarket(ctx *MarketContext) {}
