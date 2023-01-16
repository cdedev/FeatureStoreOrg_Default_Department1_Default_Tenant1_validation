// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673877901066/Conscience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Conscience

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConscienceListener is a complete listener for a parse tree produced by ConscienceParser.
type BaseConscienceListener struct{}

var _ ConscienceListener = &BaseConscienceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConscienceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConscienceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConscienceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConscienceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConscienceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConscienceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConscience is called when production conscience is entered.
func (s *BaseConscienceListener) EnterConscience(ctx *ConscienceContext) {}

// ExitConscience is called when production conscience is exited.
func (s *BaseConscienceListener) ExitConscience(ctx *ConscienceContext) {}
