// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689648490/Mean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mean

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMeanListener is a complete listener for a parse tree produced by MeanParser.
type BaseMeanListener struct{}

var _ MeanListener = &BaseMeanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMeanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMeanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMeanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMeanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMeanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMeanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMean is called when production mean is entered.
func (s *BaseMeanListener) EnterMean(ctx *MeanContext) {}

// ExitMean is called when production mean is exited.
func (s *BaseMeanListener) ExitMean(ctx *MeanContext) {}
