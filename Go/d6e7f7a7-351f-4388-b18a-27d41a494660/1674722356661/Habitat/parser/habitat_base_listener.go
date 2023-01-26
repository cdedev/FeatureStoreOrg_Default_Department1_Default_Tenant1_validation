// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722356661/Habitat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Habitat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHabitatListener is a complete listener for a parse tree produced by HabitatParser.
type BaseHabitatListener struct{}

var _ HabitatListener = &BaseHabitatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHabitatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHabitatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHabitatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHabitatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHabitatListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHabitatListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHabitat is called when production habitat is entered.
func (s *BaseHabitatListener) EnterHabitat(ctx *HabitatContext) {}

// ExitHabitat is called when production habitat is exited.
func (s *BaseHabitatListener) ExitHabitat(ctx *HabitatContext) {}
