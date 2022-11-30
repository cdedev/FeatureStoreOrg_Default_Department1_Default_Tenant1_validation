// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780393756/Valence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Valence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseValenceListener is a complete listener for a parse tree produced by ValenceParser.
type BaseValenceListener struct{}

var _ ValenceListener = &BaseValenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseValenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseValenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseValenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseValenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseValenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseValenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterValence is called when production valence is entered.
func (s *BaseValenceListener) EnterValence(ctx *ValenceContext) {}

// ExitValence is called when production valence is exited.
func (s *BaseValenceListener) ExitValence(ctx *ValenceContext) {}
