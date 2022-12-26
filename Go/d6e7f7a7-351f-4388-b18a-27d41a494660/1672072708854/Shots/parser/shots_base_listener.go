// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672072708854/Shots.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shots

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShotsListener is a complete listener for a parse tree produced by ShotsParser.
type BaseShotsListener struct{}

var _ ShotsListener = &BaseShotsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShotsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShotsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShotsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShotsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseShotsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseShotsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShots is called when production shots is entered.
func (s *BaseShotsListener) EnterShots(ctx *ShotsContext) {}

// ExitShots is called when production shots is exited.
func (s *BaseShotsListener) ExitShots(ctx *ShotsContext) {}
