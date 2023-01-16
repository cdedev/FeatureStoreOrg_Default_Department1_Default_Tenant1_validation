// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868747376/Cut.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cut

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCutListener is a complete listener for a parse tree produced by CutParser.
type BaseCutListener struct{}

var _ CutListener = &BaseCutListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCutListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCutListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCutListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCutListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCutListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCutListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCut is called when production cut is entered.
func (s *BaseCutListener) EnterCut(ctx *CutContext) {}

// ExitCut is called when production cut is exited.
func (s *BaseCutListener) ExitCut(ctx *CutContext) {}
