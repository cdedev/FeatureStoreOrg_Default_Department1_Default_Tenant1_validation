// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517657930/RadiativePower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RadiativePower

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRadiativePowerListener is a complete listener for a parse tree produced by RadiativePowerParser.
type BaseRadiativePowerListener struct{}

var _ RadiativePowerListener = &BaseRadiativePowerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRadiativePowerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRadiativePowerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRadiativePowerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRadiativePowerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRadiativePowerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRadiativePowerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRadiativepower is called when production radiativepower is entered.
func (s *BaseRadiativePowerListener) EnterRadiativepower(ctx *RadiativepowerContext) {}

// ExitRadiativepower is called when production radiativepower is exited.
func (s *BaseRadiativePowerListener) ExitRadiativepower(ctx *RadiativepowerContext) {}
