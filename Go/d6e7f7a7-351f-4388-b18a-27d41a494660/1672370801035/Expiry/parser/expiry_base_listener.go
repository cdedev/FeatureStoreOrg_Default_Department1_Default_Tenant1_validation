// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672370801035/Expiry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expiry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpiryListener is a complete listener for a parse tree produced by ExpiryParser.
type BaseExpiryListener struct{}

var _ ExpiryListener = &BaseExpiryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpiryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpiryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpiryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpiryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpiryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpiryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpiry is called when production expiry is entered.
func (s *BaseExpiryListener) EnterExpiry(ctx *ExpiryContext) {}

// ExitExpiry is called when production expiry is exited.
func (s *BaseExpiryListener) ExitExpiry(ctx *ExpiryContext) {}
