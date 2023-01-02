// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637624734/HIndex.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HIndex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHIndexListener is a complete listener for a parse tree produced by HIndexParser.
type BaseHIndexListener struct{}

var _ HIndexListener = &BaseHIndexListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHIndexListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHIndexListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHIndexListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHIndexListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHIndexListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHIndexListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHindex is called when production hindex is entered.
func (s *BaseHIndexListener) EnterHindex(ctx *HindexContext) {}

// ExitHindex is called when production hindex is exited.
func (s *BaseHIndexListener) ExitHindex(ctx *HindexContext) {}
