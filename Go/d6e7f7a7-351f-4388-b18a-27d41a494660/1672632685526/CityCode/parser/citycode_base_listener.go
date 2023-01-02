// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672632685526/CityCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CityCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCityCodeListener is a complete listener for a parse tree produced by CityCodeParser.
type BaseCityCodeListener struct{}

var _ CityCodeListener = &BaseCityCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCityCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCityCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCityCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCityCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCityCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCityCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCitycode is called when production citycode is entered.
func (s *BaseCityCodeListener) EnterCitycode(ctx *CitycodeContext) {}

// ExitCitycode is called when production citycode is exited.
func (s *BaseCityCodeListener) ExitCitycode(ctx *CitycodeContext) {}
