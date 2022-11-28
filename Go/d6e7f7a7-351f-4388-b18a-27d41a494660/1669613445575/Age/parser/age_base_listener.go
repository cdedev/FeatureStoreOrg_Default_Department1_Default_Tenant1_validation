// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669613445575/Age.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Age

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgeListener is a complete listener for a parse tree produced by AgeParser.
type BaseAgeListener struct{}

var _ AgeListener = &BaseAgeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAgeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAgeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterField0 is called when production field0 is entered.
func (s *BaseAgeListener) EnterField0(ctx *Field0Context) {}

// ExitField0 is called when production field0 is exited.
func (s *BaseAgeListener) ExitField0(ctx *Field0Context) {}
