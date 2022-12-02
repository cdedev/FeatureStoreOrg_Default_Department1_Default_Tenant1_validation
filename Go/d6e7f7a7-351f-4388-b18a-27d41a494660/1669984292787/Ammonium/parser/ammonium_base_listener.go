// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669984292787/Ammonium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ammonium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAmmoniumListener is a complete listener for a parse tree produced by AmmoniumParser.
type BaseAmmoniumListener struct{}

var _ AmmoniumListener = &BaseAmmoniumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAmmoniumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAmmoniumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAmmoniumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAmmoniumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAmmoniumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAmmoniumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAmmonium is called when production ammonium is entered.
func (s *BaseAmmoniumListener) EnterAmmonium(ctx *AmmoniumContext) {}

// ExitAmmonium is called when production ammonium is exited.
func (s *BaseAmmoniumListener) ExitAmmonium(ctx *AmmoniumContext) {}
