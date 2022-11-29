// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697681280/Acidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAcidityListener is a complete listener for a parse tree produced by AcidityParser.
type BaseAcidityListener struct{}

var _ AcidityListener = &BaseAcidityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAcidityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAcidityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAcidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAcidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAcidityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAcidityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAcidity is called when production acidity is entered.
func (s *BaseAcidityListener) EnterAcidity(ctx *AcidityContext) {}

// ExitAcidity is called when production acidity is exited.
func (s *BaseAcidityListener) ExitAcidity(ctx *AcidityContext) {}
