// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671527219193/Precipitation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Precipitation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePrecipitationListener is a complete listener for a parse tree produced by PrecipitationParser.
type BasePrecipitationListener struct{}

var _ PrecipitationListener = &BasePrecipitationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePrecipitationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePrecipitationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePrecipitationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePrecipitationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePrecipitationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePrecipitationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrecipitation is called when production precipitation is entered.
func (s *BasePrecipitationListener) EnterPrecipitation(ctx *PrecipitationContext) {}

// ExitPrecipitation is called when production precipitation is exited.
func (s *BasePrecipitationListener) ExitPrecipitation(ctx *PrecipitationContext) {}
