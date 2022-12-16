// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671189864825/Style.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Style

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStyleListener is a complete listener for a parse tree produced by StyleParser.
type BaseStyleListener struct{}

var _ StyleListener = &BaseStyleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStyleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStyleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStyleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStyleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStyleListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStyleListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStyle is called when production style is entered.
func (s *BaseStyleListener) EnterStyle(ctx *StyleContext) {}

// ExitStyle is called when production style is exited.
func (s *BaseStyleListener) ExitStyle(ctx *StyleContext) {}
