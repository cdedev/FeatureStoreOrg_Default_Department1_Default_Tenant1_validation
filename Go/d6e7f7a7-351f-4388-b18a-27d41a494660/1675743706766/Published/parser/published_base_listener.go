// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675743706766/Published.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Published

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePublishedListener is a complete listener for a parse tree produced by PublishedParser.
type BasePublishedListener struct{}

var _ PublishedListener = &BasePublishedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePublishedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePublishedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePublishedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePublishedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePublishedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePublishedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPublished is called when production published is entered.
func (s *BasePublishedListener) EnterPublished(ctx *PublishedContext) {}

// ExitPublished is called when production published is exited.
func (s *BasePublishedListener) ExitPublished(ctx *PublishedContext) {}
