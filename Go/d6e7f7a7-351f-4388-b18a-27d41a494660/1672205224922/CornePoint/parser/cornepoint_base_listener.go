// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205224922/CornePoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CornePoint

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCornePointListener is a complete listener for a parse tree produced by CornePointParser.
type BaseCornePointListener struct{}

var _ CornePointListener = &BaseCornePointListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCornePointListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCornePointListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCornePointListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCornePointListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCornePointListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCornePointListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCornepoint is called when production cornepoint is entered.
func (s *BaseCornePointListener) EnterCornepoint(ctx *CornepointContext) {}

// ExitCornepoint is called when production cornepoint is exited.
func (s *BaseCornePointListener) ExitCornepoint(ctx *CornepointContext) {}
