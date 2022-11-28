// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653223671/Slope.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Slope

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSlopeListener is a complete listener for a parse tree produced by SlopeParser.
type BaseSlopeListener struct{}

var _ SlopeListener = &BaseSlopeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSlopeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSlopeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSlopeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSlopeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSlopeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSlopeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSlope is called when production slope is entered.
func (s *BaseSlopeListener) EnterSlope(ctx *SlopeContext) {}

// ExitSlope is called when production slope is exited.
func (s *BaseSlopeListener) ExitSlope(ctx *SlopeContext) {}
