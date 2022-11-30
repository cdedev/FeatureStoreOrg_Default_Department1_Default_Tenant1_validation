// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669787921220/Barometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBarometerListener is a complete listener for a parse tree produced by BarometerParser.
type BaseBarometerListener struct{}

var _ BarometerListener = &BaseBarometerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBarometerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBarometerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBarometerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBarometerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBarometerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBarometerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBarometer is called when production barometer is entered.
func (s *BaseBarometerListener) EnterBarometer(ctx *BarometerContext) {}

// ExitBarometer is called when production barometer is exited.
func (s *BaseBarometerListener) ExitBarometer(ctx *BarometerContext) {}
