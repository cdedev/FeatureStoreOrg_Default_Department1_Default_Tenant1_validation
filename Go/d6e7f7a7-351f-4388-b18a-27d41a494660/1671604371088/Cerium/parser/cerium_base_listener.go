// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604371088/Cerium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cerium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCeriumListener is a complete listener for a parse tree produced by CeriumParser.
type BaseCeriumListener struct{}

var _ CeriumListener = &BaseCeriumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCeriumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCeriumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCeriumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCeriumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCeriumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCeriumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCerium is called when production cerium is entered.
func (s *BaseCeriumListener) EnterCerium(ctx *CeriumContext) {}

// ExitCerium is called when production cerium is exited.
func (s *BaseCeriumListener) ExitCerium(ctx *CeriumContext) {}
