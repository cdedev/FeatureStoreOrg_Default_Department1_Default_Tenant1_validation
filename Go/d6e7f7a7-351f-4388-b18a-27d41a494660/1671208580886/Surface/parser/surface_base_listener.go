// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671208580886/Surface.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Surface

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSurfaceListener is a complete listener for a parse tree produced by SurfaceParser.
type BaseSurfaceListener struct{}

var _ SurfaceListener = &BaseSurfaceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSurfaceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSurfaceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSurfaceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSurfaceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSurfaceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSurfaceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSurface is called when production surface is entered.
func (s *BaseSurfaceListener) EnterSurface(ctx *SurfaceContext) {}

// ExitSurface is called when production surface is exited.
func (s *BaseSurfaceListener) ExitSurface(ctx *SurfaceContext) {}
