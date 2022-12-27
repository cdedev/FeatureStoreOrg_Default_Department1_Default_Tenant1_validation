// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118835736/RideTime.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RideTime

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRideTimeListener is a complete listener for a parse tree produced by RideTimeParser.
type BaseRideTimeListener struct{}

var _ RideTimeListener = &BaseRideTimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRideTimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRideTimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRideTimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRideTimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRideTimeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRideTimeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRidetime is called when production ridetime is entered.
func (s *BaseRideTimeListener) EnterRidetime(ctx *RidetimeContext) {}

// ExitRidetime is called when production ridetime is exited.
func (s *BaseRideTimeListener) ExitRidetime(ctx *RidetimeContext) {}
