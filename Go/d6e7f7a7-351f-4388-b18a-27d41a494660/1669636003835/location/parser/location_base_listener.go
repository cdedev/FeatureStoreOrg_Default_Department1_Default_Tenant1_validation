// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669636003835/location.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // location

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselocationListener is a complete listener for a parse tree produced by locationParser.
type BaselocationListener struct{}

var _ locationListener = &BaselocationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselocationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselocationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselocationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselocationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaselocationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaselocationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLocation is called when production location is entered.
func (s *BaselocationListener) EnterLocation(ctx *LocationContext) {}

// ExitLocation is called when production location is exited.
func (s *BaselocationListener) ExitLocation(ctx *LocationContext) {}
