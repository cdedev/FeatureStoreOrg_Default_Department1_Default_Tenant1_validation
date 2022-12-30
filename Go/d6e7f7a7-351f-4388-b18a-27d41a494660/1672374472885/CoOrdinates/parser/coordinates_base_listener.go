// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672374472885/CoOrdinates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CoOrdinates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCoOrdinatesListener is a complete listener for a parse tree produced by CoOrdinatesParser.
type BaseCoOrdinatesListener struct{}

var _ CoOrdinatesListener = &BaseCoOrdinatesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCoOrdinatesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCoOrdinatesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCoOrdinatesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCoOrdinatesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCoOrdinatesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCoOrdinatesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCoordinates is called when production coordinates is entered.
func (s *BaseCoOrdinatesListener) EnterCoordinates(ctx *CoordinatesContext) {}

// ExitCoordinates is called when production coordinates is exited.
func (s *BaseCoOrdinatesListener) ExitCoordinates(ctx *CoordinatesContext) {}
