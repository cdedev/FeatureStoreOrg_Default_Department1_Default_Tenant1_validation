// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672380142110/Bleeding.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bleeding

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBleedingListener is a complete listener for a parse tree produced by BleedingParser.
type BaseBleedingListener struct{}

var _ BleedingListener = &BaseBleedingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBleedingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBleedingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBleedingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBleedingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBleedingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBleedingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBleeding is called when production bleeding is entered.
func (s *BaseBleedingListener) EnterBleeding(ctx *BleedingContext) {}

// ExitBleeding is called when production bleeding is exited.
func (s *BaseBleedingListener) ExitBleeding(ctx *BleedingContext) {}
