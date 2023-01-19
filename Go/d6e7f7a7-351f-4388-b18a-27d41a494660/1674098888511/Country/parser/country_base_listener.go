// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674098888511/Country.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Country

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCountryListener is a complete listener for a parse tree produced by CountryParser.
type BaseCountryListener struct{}

var _ CountryListener = &BaseCountryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCountryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCountryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCountryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCountryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCountryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCountryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCountry is called when production country is entered.
func (s *BaseCountryListener) EnterCountry(ctx *CountryContext) {}

// ExitCountry is called when production country is exited.
func (s *BaseCountryListener) ExitCountry(ctx *CountryContext) {}
