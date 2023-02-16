// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524772055/Popularity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Popularity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePopularityListener is a complete listener for a parse tree produced by PopularityParser.
type BasePopularityListener struct{}

var _ PopularityListener = &BasePopularityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePopularityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePopularityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePopularityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePopularityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePopularityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePopularityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPopularity is called when production popularity is entered.
func (s *BasePopularityListener) EnterPopularity(ctx *PopularityContext) {}

// ExitPopularity is called when production popularity is exited.
func (s *BasePopularityListener) ExitPopularity(ctx *PopularityContext) {}
