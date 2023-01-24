// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538873750/Noise.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Noise

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNoiseListener is a complete listener for a parse tree produced by NoiseParser.
type BaseNoiseListener struct{}

var _ NoiseListener = &BaseNoiseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNoiseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNoiseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNoiseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNoiseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNoiseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNoiseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNoise is called when production noise is entered.
func (s *BaseNoiseListener) EnterNoise(ctx *NoiseContext) {}

// ExitNoise is called when production noise is exited.
func (s *BaseNoiseListener) ExitNoise(ctx *NoiseContext) {}
