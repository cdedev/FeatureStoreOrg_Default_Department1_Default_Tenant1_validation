// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238830271/Oily.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oily

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOilyListener is a complete listener for a parse tree produced by OilyParser.
type BaseOilyListener struct{}

var _ OilyListener = &BaseOilyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOilyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOilyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOilyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOilyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOilyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOilyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOily is called when production oily is entered.
func (s *BaseOilyListener) EnterOily(ctx *OilyContext) {}

// ExitOily is called when production oily is exited.
func (s *BaseOilyListener) ExitOily(ctx *OilyContext) {}
