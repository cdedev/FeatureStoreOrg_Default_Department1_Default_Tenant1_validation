// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121411533/Rabbits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rabbits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRabbitsListener is a complete listener for a parse tree produced by RabbitsParser.
type BaseRabbitsListener struct{}

var _ RabbitsListener = &BaseRabbitsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRabbitsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRabbitsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRabbitsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRabbitsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRabbitsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRabbitsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRabbits is called when production rabbits is entered.
func (s *BaseRabbitsListener) EnterRabbits(ctx *RabbitsContext) {}

// ExitRabbits is called when production rabbits is exited.
func (s *BaseRabbitsListener) ExitRabbits(ctx *RabbitsContext) {}
