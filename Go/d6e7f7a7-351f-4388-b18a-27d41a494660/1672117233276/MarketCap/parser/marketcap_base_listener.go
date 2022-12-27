// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672117233276/MarketCap.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarketCap

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarketCapListener is a complete listener for a parse tree produced by MarketCapParser.
type BaseMarketCapListener struct{}

var _ MarketCapListener = &BaseMarketCapListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarketCapListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarketCapListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarketCapListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarketCapListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarketCapListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarketCapListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarketcap is called when production marketcap is entered.
func (s *BaseMarketCapListener) EnterMarketcap(ctx *MarketcapContext) {}

// ExitMarketcap is called when production marketcap is exited.
func (s *BaseMarketCapListener) ExitMarketcap(ctx *MarketcapContext) {}
