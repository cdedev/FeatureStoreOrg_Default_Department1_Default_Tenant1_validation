// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538797844/Loud.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loud

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLoudListener is a complete listener for a parse tree produced by LoudParser.
type BaseLoudListener struct{}

var _ LoudListener = &BaseLoudListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLoudListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLoudListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLoudListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLoudListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLoudListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLoudListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLoud is called when production loud is entered.
func (s *BaseLoudListener) EnterLoud(ctx *LoudContext) {}

// ExitLoud is called when production loud is exited.
func (s *BaseLoudListener) ExitLoud(ctx *LoudContext) {}
