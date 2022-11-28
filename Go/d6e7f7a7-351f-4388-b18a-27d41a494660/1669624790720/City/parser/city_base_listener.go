// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624790720/City.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // City

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCityListener is a complete listener for a parse tree produced by CityParser.
type BaseCityListener struct{}

var _ CityListener = &BaseCityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCity is called when production city is entered.
func (s *BaseCityListener) EnterCity(ctx *CityContext) {}

// ExitCity is called when production city is exited.
func (s *BaseCityListener) ExitCity(ctx *CityContext) {}
