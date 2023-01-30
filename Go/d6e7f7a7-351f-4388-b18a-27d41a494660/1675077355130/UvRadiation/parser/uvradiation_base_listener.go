// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675077355130/UvRadiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // UvRadiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUvRadiationListener is a complete listener for a parse tree produced by UvRadiationParser.
type BaseUvRadiationListener struct{}

var _ UvRadiationListener = &BaseUvRadiationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUvRadiationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUvRadiationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUvRadiationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUvRadiationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUvRadiationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUvRadiationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUvradiation is called when production uvradiation is entered.
func (s *BaseUvRadiationListener) EnterUvradiation(ctx *UvradiationContext) {}

// ExitUvradiation is called when production uvradiation is exited.
func (s *BaseUvRadiationListener) ExitUvradiation(ctx *UvradiationContext) {}
