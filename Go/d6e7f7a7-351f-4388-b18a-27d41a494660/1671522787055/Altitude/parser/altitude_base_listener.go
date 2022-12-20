// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671522787055/Altitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Altitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAltitudeListener is a complete listener for a parse tree produced by AltitudeParser.
type BaseAltitudeListener struct{}

var _ AltitudeListener = &BaseAltitudeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAltitudeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAltitudeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAltitudeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAltitudeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAltitudeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAltitudeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAltitude is called when production altitude is entered.
func (s *BaseAltitudeListener) EnterAltitude(ctx *AltitudeContext) {}

// ExitAltitude is called when production altitude is exited.
func (s *BaseAltitudeListener) ExitAltitude(ctx *AltitudeContext) {}
