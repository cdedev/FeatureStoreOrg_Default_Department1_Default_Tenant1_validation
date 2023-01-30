// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675076839293/Sunrise.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sunrise

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSunriseListener is a complete listener for a parse tree produced by SunriseParser.
type BaseSunriseListener struct{}

var _ SunriseListener = &BaseSunriseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSunriseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSunriseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSunriseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSunriseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSunriseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSunriseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSunrise is called when production sunrise is entered.
func (s *BaseSunriseListener) EnterSunrise(ctx *SunriseContext) {}

// ExitSunrise is called when production sunrise is exited.
func (s *BaseSunriseListener) ExitSunrise(ctx *SunriseContext) {}
