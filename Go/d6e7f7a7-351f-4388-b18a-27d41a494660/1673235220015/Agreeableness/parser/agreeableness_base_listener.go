// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235220015/Agreeableness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreeableness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgreeablenessListener is a complete listener for a parse tree produced by AgreeablenessParser.
type BaseAgreeablenessListener struct{}

var _ AgreeablenessListener = &BaseAgreeablenessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgreeablenessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgreeablenessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgreeablenessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgreeablenessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAgreeablenessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAgreeablenessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAgreeableness is called when production agreeableness is entered.
func (s *BaseAgreeablenessListener) EnterAgreeableness(ctx *AgreeablenessContext) {}

// ExitAgreeableness is called when production agreeableness is exited.
func (s *BaseAgreeablenessListener) ExitAgreeableness(ctx *AgreeablenessContext) {}
