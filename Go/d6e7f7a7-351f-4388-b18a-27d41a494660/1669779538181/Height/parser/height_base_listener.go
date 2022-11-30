// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669779538181/Height.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Height

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHeightListener is a complete listener for a parse tree produced by HeightParser.
type BaseHeightListener struct{}

var _ HeightListener = &BaseHeightListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHeightListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHeightListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHeightListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHeightListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHeightListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHeightListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHeight is called when production height is entered.
func (s *BaseHeightListener) EnterHeight(ctx *HeightContext) {}

// ExitHeight is called when production height is exited.
func (s *BaseHeightListener) ExitHeight(ctx *HeightContext) {}
