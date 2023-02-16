// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524720628/Loudness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loudness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLoudnessListener is a complete listener for a parse tree produced by LoudnessParser.
type BaseLoudnessListener struct{}

var _ LoudnessListener = &BaseLoudnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLoudnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLoudnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLoudnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLoudnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLoudnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLoudnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLoudness is called when production loudness is entered.
func (s *BaseLoudnessListener) EnterLoudness(ctx *LoudnessContext) {}

// ExitLoudness is called when production loudness is exited.
func (s *BaseLoudnessListener) ExitLoudness(ctx *LoudnessContext) {}
