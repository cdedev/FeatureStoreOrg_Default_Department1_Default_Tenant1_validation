// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690493750/Sodium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sodium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSodiumListener is a complete listener for a parse tree produced by SodiumParser.
type BaseSodiumListener struct{}

var _ SodiumListener = &BaseSodiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSodiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSodiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSodiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSodiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSodiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSodiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSodium is called when production sodium is entered.
func (s *BaseSodiumListener) EnterSodium(ctx *SodiumContext) {}

// ExitSodium is called when production sodium is exited.
func (s *BaseSodiumListener) ExitSodium(ctx *SodiumContext) {}
