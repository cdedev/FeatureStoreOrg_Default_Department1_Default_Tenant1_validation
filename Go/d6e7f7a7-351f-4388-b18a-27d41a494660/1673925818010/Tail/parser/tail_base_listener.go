// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925818010/Tail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tail

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTailListener is a complete listener for a parse tree produced by TailParser.
type BaseTailListener struct{}

var _ TailListener = &BaseTailListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTailListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTailListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTailListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTailListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTailListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTailListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTail is called when production tail is entered.
func (s *BaseTailListener) EnterTail(ctx *TailContext) {}

// ExitTail is called when production tail is exited.
func (s *BaseTailListener) ExitTail(ctx *TailContext) {}
