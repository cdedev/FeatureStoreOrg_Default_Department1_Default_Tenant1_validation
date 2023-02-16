// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522774644/Publisher.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Publisher

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePublisherListener is a complete listener for a parse tree produced by PublisherParser.
type BasePublisherListener struct{}

var _ PublisherListener = &BasePublisherListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePublisherListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePublisherListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePublisherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePublisherListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePublisherListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePublisherListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPublisher is called when production publisher is entered.
func (s *BasePublisherListener) EnterPublisher(ctx *PublisherContext) {}

// ExitPublisher is called when production publisher is exited.
func (s *BasePublisherListener) ExitPublisher(ctx *PublisherContext) {}
