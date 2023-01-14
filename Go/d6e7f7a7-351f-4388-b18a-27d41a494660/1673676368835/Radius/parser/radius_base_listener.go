// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676368835/Radius.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radius

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRadiusListener is a complete listener for a parse tree produced by RadiusParser.
type BaseRadiusListener struct{}

var _ RadiusListener = &BaseRadiusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRadiusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRadiusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRadiusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRadiusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRadiusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRadiusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRadius is called when production radius is entered.
func (s *BaseRadiusListener) EnterRadius(ctx *RadiusContext) {}

// ExitRadius is called when production radius is exited.
func (s *BaseRadiusListener) ExitRadius(ctx *RadiusContext) {}
