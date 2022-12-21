// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603867381/Bismuth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bismuth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBismuthListener is a complete listener for a parse tree produced by BismuthParser.
type BaseBismuthListener struct{}

var _ BismuthListener = &BaseBismuthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBismuthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBismuthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBismuthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBismuthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBismuthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBismuthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBismuth is called when production bismuth is entered.
func (s *BaseBismuthListener) EnterBismuth(ctx *BismuthContext) {}

// ExitBismuth is called when production bismuth is exited.
func (s *BaseBismuthListener) ExitBismuth(ctx *BismuthContext) {}
