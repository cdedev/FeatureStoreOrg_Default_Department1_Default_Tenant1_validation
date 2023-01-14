// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673670414326/WeatherType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeatherTypeListener is a complete listener for a parse tree produced by WeatherTypeParser.
type BaseWeatherTypeListener struct{}

var _ WeatherTypeListener = &BaseWeatherTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeatherTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeatherTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeatherTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeatherTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeatherTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeatherTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeathertype is called when production weathertype is entered.
func (s *BaseWeatherTypeListener) EnterWeathertype(ctx *WeathertypeContext) {}

// ExitWeathertype is called when production weathertype is exited.
func (s *BaseWeatherTypeListener) ExitWeathertype(ctx *WeathertypeContext) {}
