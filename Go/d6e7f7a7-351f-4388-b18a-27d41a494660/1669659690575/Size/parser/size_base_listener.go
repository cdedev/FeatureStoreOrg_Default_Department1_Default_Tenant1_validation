// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659690575/Size.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Size

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSizeListener is a complete listener for a parse tree produced by SizeParser.
type BaseSizeListener struct{}

var _ SizeListener = &BaseSizeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSizeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSizeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSizeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSizeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSizeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSizeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSize is called when production size is entered.
func (s *BaseSizeListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BaseSizeListener) ExitSize(ctx *SizeContext) {}
