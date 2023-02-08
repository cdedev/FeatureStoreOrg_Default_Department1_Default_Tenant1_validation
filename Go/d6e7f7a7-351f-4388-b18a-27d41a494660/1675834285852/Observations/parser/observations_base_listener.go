// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834285852/Observations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Observations

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObservationsListener is a complete listener for a parse tree produced by ObservationsParser.
type BaseObservationsListener struct{}

var _ ObservationsListener = &BaseObservationsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObservationsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObservationsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObservationsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObservationsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseObservationsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseObservationsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterObservations is called when production observations is entered.
func (s *BaseObservationsListener) EnterObservations(ctx *ObservationsContext) {}

// ExitObservations is called when production observations is exited.
func (s *BaseObservationsListener) ExitObservations(ctx *ObservationsContext) {}
