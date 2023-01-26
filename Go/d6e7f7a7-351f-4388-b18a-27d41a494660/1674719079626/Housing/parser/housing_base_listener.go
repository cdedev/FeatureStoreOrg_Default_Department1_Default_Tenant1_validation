// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719079626/Housing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Housing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHousingListener is a complete listener for a parse tree produced by HousingParser.
type BaseHousingListener struct{}

var _ HousingListener = &BaseHousingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHousingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHousingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHousingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHousingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHousingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHousingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHousing is called when production housing is entered.
func (s *BaseHousingListener) EnterHousing(ctx *HousingContext) {}

// ExitHousing is called when production housing is exited.
func (s *BaseHousingListener) ExitHousing(ctx *HousingContext) {}
