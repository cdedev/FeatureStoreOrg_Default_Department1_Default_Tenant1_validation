// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628361272/Humidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Humidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHumidityListener is a complete listener for a parse tree produced by HumidityParser.
type BaseHumidityListener struct{}

var _ HumidityListener = &BaseHumidityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHumidityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHumidityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHumidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHumidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHumidityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHumidityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHumidity is called when production humidity is entered.
func (s *BaseHumidityListener) EnterHumidity(ctx *HumidityContext) {}

// ExitHumidity is called when production humidity is exited.
func (s *BaseHumidityListener) ExitHumidity(ctx *HumidityContext) {}
