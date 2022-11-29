// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699291398/Minoraxis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Minoraxis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMinoraxisListener is a complete listener for a parse tree produced by MinoraxisParser.
type BaseMinoraxisListener struct{}

var _ MinoraxisListener = &BaseMinoraxisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinoraxisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinoraxisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinoraxisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinoraxisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMinoraxisListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMinoraxisListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMinoraxis is called when production minoraxis is entered.
func (s *BaseMinoraxisListener) EnterMinoraxis(ctx *MinoraxisContext) {}

// ExitMinoraxis is called when production minoraxis is exited.
func (s *BaseMinoraxisListener) ExitMinoraxis(ctx *MinoraxisContext) {}
