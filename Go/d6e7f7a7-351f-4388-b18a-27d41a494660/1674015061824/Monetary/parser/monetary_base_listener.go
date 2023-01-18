// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674015061824/Monetary.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Monetary

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMonetaryListener is a complete listener for a parse tree produced by MonetaryParser.
type BaseMonetaryListener struct{}

var _ MonetaryListener = &BaseMonetaryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMonetaryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMonetaryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMonetaryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMonetaryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMonetaryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMonetaryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMonetary is called when production monetary is entered.
func (s *BaseMonetaryListener) EnterMonetary(ctx *MonetaryContext) {}

// ExitMonetary is called when production monetary is exited.
func (s *BaseMonetaryListener) ExitMonetary(ctx *MonetaryContext) {}
