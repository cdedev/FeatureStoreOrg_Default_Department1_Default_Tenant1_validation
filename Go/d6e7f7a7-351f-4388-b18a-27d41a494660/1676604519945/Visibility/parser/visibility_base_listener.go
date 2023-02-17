// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676604519945/Visibility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visibility

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVisibilityListener is a complete listener for a parse tree produced by VisibilityParser.
type BaseVisibilityListener struct{}

var _ VisibilityListener = &BaseVisibilityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVisibilityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVisibilityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVisibilityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVisibilityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVisibilityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVisibilityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseVisibilityListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseVisibilityListener) ExitVisibility(ctx *VisibilityContext) {}
