// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676512220/Symmetry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Symmetry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSymmetryListener is a complete listener for a parse tree produced by SymmetryParser.
type BaseSymmetryListener struct{}

var _ SymmetryListener = &BaseSymmetryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSymmetryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSymmetryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSymmetryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSymmetryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSymmetryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSymmetryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSymmetry is called when production symmetry is entered.
func (s *BaseSymmetryListener) EnterSymmetry(ctx *SymmetryContext) {}

// ExitSymmetry is called when production symmetry is exited.
func (s *BaseSymmetryListener) ExitSymmetry(ctx *SymmetryContext) {}
