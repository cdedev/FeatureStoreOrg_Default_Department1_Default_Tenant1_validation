// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693230058/Kilometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kilometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKilometerListener is a complete listener for a parse tree produced by KilometerParser.
type BaseKilometerListener struct{}

var _ KilometerListener = &BaseKilometerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKilometerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKilometerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKilometerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKilometerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKilometerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKilometerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKilometer is called when production kilometer is entered.
func (s *BaseKilometerListener) EnterKilometer(ctx *KilometerContext) {}

// ExitKilometer is called when production kilometer is exited.
func (s *BaseKilometerListener) ExitKilometer(ctx *KilometerContext) {}
