// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972227227/Low.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Low

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLowListener is a complete listener for a parse tree produced by LowParser.
type BaseLowListener struct{}

var _ LowListener = &BaseLowListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLowListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLowListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLowListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLowListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLowListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLowListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLow is called when production low is entered.
func (s *BaseLowListener) EnterLow(ctx *LowContext) {}

// ExitLow is called when production low is exited.
func (s *BaseLowListener) ExitLow(ctx *LowContext) {}
