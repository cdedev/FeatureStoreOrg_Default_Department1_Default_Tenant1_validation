// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118968290/RideLength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RideLength

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRideLengthListener is a complete listener for a parse tree produced by RideLengthParser.
type BaseRideLengthListener struct{}

var _ RideLengthListener = &BaseRideLengthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRideLengthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRideLengthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRideLengthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRideLengthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRideLengthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRideLengthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRidelength is called when production ridelength is entered.
func (s *BaseRideLengthListener) EnterRidelength(ctx *RidelengthContext) {}

// ExitRidelength is called when production ridelength is exited.
func (s *BaseRideLengthListener) ExitRidelength(ctx *RidelengthContext) {}
