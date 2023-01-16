// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673877803247/Ataxia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ataxia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAtaxiaListener is a complete listener for a parse tree produced by AtaxiaParser.
type BaseAtaxiaListener struct{}

var _ AtaxiaListener = &BaseAtaxiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAtaxiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAtaxiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAtaxiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAtaxiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAtaxiaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAtaxiaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAtaxia is called when production ataxia is entered.
func (s *BaseAtaxiaListener) EnterAtaxia(ctx *AtaxiaContext) {}

// ExitAtaxia is called when production ataxia is exited.
func (s *BaseAtaxiaListener) ExitAtaxia(ctx *AtaxiaContext) {}
