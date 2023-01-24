// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532608460/Margin.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Margin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarginListener is a complete listener for a parse tree produced by MarginParser.
type BaseMarginListener struct{}

var _ MarginListener = &BaseMarginListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarginListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarginListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarginListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarginListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarginListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarginListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMargin is called when production margin is entered.
func (s *BaseMarginListener) EnterMargin(ctx *MarginContext) {}

// ExitMargin is called when production margin is exited.
func (s *BaseMarginListener) ExitMargin(ctx *MarginContext) {}
