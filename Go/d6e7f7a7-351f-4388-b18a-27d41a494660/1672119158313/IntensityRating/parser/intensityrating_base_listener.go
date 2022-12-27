// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119158313/IntensityRating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntensityRating

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntensityRatingListener is a complete listener for a parse tree produced by IntensityRatingParser.
type BaseIntensityRatingListener struct{}

var _ IntensityRatingListener = &BaseIntensityRatingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntensityRatingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntensityRatingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntensityRatingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntensityRatingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntensityRatingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntensityRatingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIntensityrating is called when production intensityrating is entered.
func (s *BaseIntensityRatingListener) EnterIntensityrating(ctx *IntensityratingContext) {}

// ExitIntensityrating is called when production intensityrating is exited.
func (s *BaseIntensityRatingListener) ExitIntensityrating(ctx *IntensityratingContext) {}
