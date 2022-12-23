// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771286494/WeatherTemperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherTemperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeatherTemperatureListener is a complete listener for a parse tree produced by WeatherTemperatureParser.
type BaseWeatherTemperatureListener struct{}

var _ WeatherTemperatureListener = &BaseWeatherTemperatureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeatherTemperatureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeatherTemperatureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeatherTemperatureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeatherTemperatureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeatherTemperatureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeatherTemperatureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeathertemperature is called when production weathertemperature is entered.
func (s *BaseWeatherTemperatureListener) EnterWeathertemperature(ctx *WeathertemperatureContext) {}

// ExitWeathertemperature is called when production weathertemperature is exited.
func (s *BaseWeatherTemperatureListener) ExitWeathertemperature(ctx *WeathertemperatureContext) {}
