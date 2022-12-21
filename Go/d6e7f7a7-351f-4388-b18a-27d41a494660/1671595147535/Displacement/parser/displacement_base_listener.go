// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671595147535/Displacement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Displacement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDisplacementListener is a complete listener for a parse tree produced by DisplacementParser.
type BaseDisplacementListener struct{}

var _ DisplacementListener = &BaseDisplacementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDisplacementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDisplacementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDisplacementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDisplacementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDisplacementListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDisplacementListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDisplacement is called when production displacement is entered.
func (s *BaseDisplacementListener) EnterDisplacement(ctx *DisplacementContext) {}

// ExitDisplacement is called when production displacement is exited.
func (s *BaseDisplacementListener) ExitDisplacement(ctx *DisplacementContext) {}
