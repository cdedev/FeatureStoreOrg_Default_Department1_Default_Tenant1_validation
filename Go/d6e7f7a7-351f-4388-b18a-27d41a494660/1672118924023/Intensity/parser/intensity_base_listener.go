// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118924023/Intensity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Intensity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntensityListener is a complete listener for a parse tree produced by IntensityParser.
type BaseIntensityListener struct{}

var _ IntensityListener = &BaseIntensityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntensityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntensityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntensityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntensityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntensityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntensityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIntensity is called when production intensity is entered.
func (s *BaseIntensityListener) EnterIntensity(ctx *IntensityContext) {}

// ExitIntensity is called when production intensity is exited.
func (s *BaseIntensityListener) ExitIntensity(ctx *IntensityContext) {}
