// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656626453/Airquality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airquality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAirqualityListener is a complete listener for a parse tree produced by AirqualityParser.
type BaseAirqualityListener struct{}

var _ AirqualityListener = &BaseAirqualityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAirqualityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAirqualityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAirqualityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAirqualityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAirqualityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAirqualityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAirquality is called when production airquality is entered.
func (s *BaseAirqualityListener) EnterAirquality(ctx *AirqualityContext) {}

// ExitAirquality is called when production airquality is exited.
func (s *BaseAirqualityListener) ExitAirquality(ctx *AirqualityContext) {}
