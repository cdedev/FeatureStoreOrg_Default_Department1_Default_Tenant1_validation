// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656997157/Weather.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weather

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeatherListener is a complete listener for a parse tree produced by WeatherParser.
type BaseWeatherListener struct{}

var _ WeatherListener = &BaseWeatherListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeatherListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeatherListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeatherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeatherListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeatherListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeatherListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeather is called when production weather is entered.
func (s *BaseWeatherListener) EnterWeather(ctx *WeatherContext) {}

// ExitWeather is called when production weather is exited.
func (s *BaseWeatherListener) ExitWeather(ctx *WeatherContext) {}
