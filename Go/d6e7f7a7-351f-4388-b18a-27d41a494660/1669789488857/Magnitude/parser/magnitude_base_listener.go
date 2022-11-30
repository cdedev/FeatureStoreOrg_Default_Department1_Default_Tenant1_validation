// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789488857/Magnitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Magnitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMagnitudeListener is a complete listener for a parse tree produced by MagnitudeParser.
type BaseMagnitudeListener struct{}

var _ MagnitudeListener = &BaseMagnitudeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMagnitudeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMagnitudeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMagnitudeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMagnitudeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMagnitudeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMagnitudeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMagnitude is called when production magnitude is entered.
func (s *BaseMagnitudeListener) EnterMagnitude(ctx *MagnitudeContext) {}

// ExitMagnitude is called when production magnitude is exited.
func (s *BaseMagnitudeListener) ExitMagnitude(ctx *MagnitudeContext) {}
