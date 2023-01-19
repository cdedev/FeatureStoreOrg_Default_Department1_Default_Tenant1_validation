// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674104270059/GlobalRadiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GlobalRadiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGlobalRadiationListener is a complete listener for a parse tree produced by GlobalRadiationParser.
type BaseGlobalRadiationListener struct{}

var _ GlobalRadiationListener = &BaseGlobalRadiationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGlobalRadiationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGlobalRadiationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGlobalRadiationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGlobalRadiationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGlobalRadiationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGlobalRadiationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGlobalradiation is called when production globalradiation is entered.
func (s *BaseGlobalRadiationListener) EnterGlobalradiation(ctx *GlobalradiationContext) {}

// ExitGlobalradiation is called when production globalradiation is exited.
func (s *BaseGlobalRadiationListener) ExitGlobalradiation(ctx *GlobalradiationContext) {}
