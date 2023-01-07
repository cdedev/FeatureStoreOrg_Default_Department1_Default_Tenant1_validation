// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083556717/Perpetrator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Perpetrator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePerpetratorListener is a complete listener for a parse tree produced by PerpetratorParser.
type BasePerpetratorListener struct{}

var _ PerpetratorListener = &BasePerpetratorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePerpetratorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePerpetratorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePerpetratorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePerpetratorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePerpetratorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePerpetratorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPerpetrator is called when production perpetrator is entered.
func (s *BasePerpetratorListener) EnterPerpetrator(ctx *PerpetratorContext) {}

// ExitPerpetrator is called when production perpetrator is exited.
func (s *BasePerpetratorListener) ExitPerpetrator(ctx *PerpetratorContext) {}
