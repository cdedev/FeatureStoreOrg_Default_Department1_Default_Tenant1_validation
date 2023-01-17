// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924843070/Airborne.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airborne

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAirborneListener is a complete listener for a parse tree produced by AirborneParser.
type BaseAirborneListener struct{}

var _ AirborneListener = &BaseAirborneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAirborneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAirborneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAirborneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAirborneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAirborneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAirborneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAirborne is called when production airborne is entered.
func (s *BaseAirborneListener) EnterAirborne(ctx *AirborneContext) {}

// ExitAirborne is called when production airborne is exited.
func (s *BaseAirborneListener) ExitAirborne(ctx *AirborneContext) {}
