// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925871862/Toothed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Toothed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseToothedListener is a complete listener for a parse tree produced by ToothedParser.
type BaseToothedListener struct{}

var _ ToothedListener = &BaseToothedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseToothedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseToothedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseToothedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseToothedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseToothedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseToothedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterToothed is called when production toothed is entered.
func (s *BaseToothedListener) EnterToothed(ctx *ToothedContext) {}

// ExitToothed is called when production toothed is exited.
func (s *BaseToothedListener) ExitToothed(ctx *ToothedContext) {}
