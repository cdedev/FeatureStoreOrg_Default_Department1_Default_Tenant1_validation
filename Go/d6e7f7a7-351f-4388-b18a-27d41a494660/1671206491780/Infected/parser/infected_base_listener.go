// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206491780/Infected.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Infected

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInfectedListener is a complete listener for a parse tree produced by InfectedParser.
type BaseInfectedListener struct{}

var _ InfectedListener = &BaseInfectedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInfectedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInfectedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInfectedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInfectedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInfectedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInfectedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInfected is called when production infected is entered.
func (s *BaseInfectedListener) EnterInfected(ctx *InfectedContext) {}

// ExitInfected is called when production infected is exited.
func (s *BaseInfectedListener) ExitInfected(ctx *InfectedContext) {}
