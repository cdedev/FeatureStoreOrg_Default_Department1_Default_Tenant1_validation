// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096810128/Pitch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pitch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePitchListener is a complete listener for a parse tree produced by PitchParser.
type BasePitchListener struct{}

var _ PitchListener = &BasePitchListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePitchListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePitchListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePitchListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePitchListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePitchListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePitchListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPitch is called when production pitch is entered.
func (s *BasePitchListener) EnterPitch(ctx *PitchContext) {}

// ExitPitch is called when production pitch is exited.
func (s *BasePitchListener) ExitPitch(ctx *PitchContext) {}
