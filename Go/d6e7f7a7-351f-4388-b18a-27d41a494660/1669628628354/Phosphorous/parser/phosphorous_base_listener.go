// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628628354/Phosphorous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phosphorous

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhosphorousListener is a complete listener for a parse tree produced by PhosphorousParser.
type BasePhosphorousListener struct{}

var _ PhosphorousListener = &BasePhosphorousListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhosphorousListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhosphorousListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhosphorousListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhosphorousListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhosphorousListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhosphorousListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhosphorous is called when production phosphorous is entered.
func (s *BasePhosphorousListener) EnterPhosphorous(ctx *PhosphorousContext) {}

// ExitPhosphorous is called when production phosphorous is exited.
func (s *BasePhosphorousListener) ExitPhosphorous(ctx *PhosphorousContext) {}
