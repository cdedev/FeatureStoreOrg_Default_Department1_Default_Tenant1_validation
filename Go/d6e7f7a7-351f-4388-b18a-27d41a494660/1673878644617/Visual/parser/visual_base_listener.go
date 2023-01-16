// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878644617/Visual.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visual

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVisualListener is a complete listener for a parse tree produced by VisualParser.
type BaseVisualListener struct{}

var _ VisualListener = &BaseVisualListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVisualListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVisualListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVisualListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVisualListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVisualListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVisualListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVisual is called when production visual is entered.
func (s *BaseVisualListener) EnterVisual(ctx *VisualContext) {}

// ExitVisual is called when production visual is exited.
func (s *BaseVisualListener) ExitVisual(ctx *VisualContext) {}
