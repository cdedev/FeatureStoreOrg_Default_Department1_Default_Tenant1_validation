// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673243876019/Radiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRadiationListener is a complete listener for a parse tree produced by RadiationParser.
type BaseRadiationListener struct{}

var _ RadiationListener = &BaseRadiationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRadiationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRadiationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRadiationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRadiationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRadiationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRadiationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRadiation is called when production radiation is entered.
func (s *BaseRadiationListener) EnterRadiation(ctx *RadiationContext) {}

// ExitRadiation is called when production radiation is exited.
func (s *BaseRadiationListener) ExitRadiation(ctx *RadiationContext) {}
