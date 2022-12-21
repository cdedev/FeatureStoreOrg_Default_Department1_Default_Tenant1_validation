// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603693191/Argon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Argon

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArgonListener is a complete listener for a parse tree produced by ArgonParser.
type BaseArgonListener struct{}

var _ ArgonListener = &BaseArgonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArgonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArgonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArgonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArgonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseArgonListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseArgonListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArgon is called when production argon is entered.
func (s *BaseArgonListener) EnterArgon(ctx *ArgonContext) {}

// ExitArgon is called when production argon is exited.
func (s *BaseArgonListener) ExitArgon(ctx *ArgonContext) {}
