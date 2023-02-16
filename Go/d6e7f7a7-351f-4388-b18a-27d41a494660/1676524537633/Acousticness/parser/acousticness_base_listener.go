// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524537633/Acousticness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acousticness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAcousticnessListener is a complete listener for a parse tree produced by AcousticnessParser.
type BaseAcousticnessListener struct{}

var _ AcousticnessListener = &BaseAcousticnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAcousticnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAcousticnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAcousticnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAcousticnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAcousticnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAcousticnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAcousticness is called when production acousticness is entered.
func (s *BaseAcousticnessListener) EnterAcousticness(ctx *AcousticnessContext) {}

// ExitAcousticness is called when production acousticness is exited.
func (s *BaseAcousticnessListener) ExitAcousticness(ctx *AcousticnessContext) {}
