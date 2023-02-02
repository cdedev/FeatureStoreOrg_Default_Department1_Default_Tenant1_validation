// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675316683481/Species.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Species

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSpeciesListener is a complete listener for a parse tree produced by SpeciesParser.
type BaseSpeciesListener struct{}

var _ SpeciesListener = &BaseSpeciesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSpeciesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSpeciesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSpeciesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSpeciesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSpeciesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSpeciesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSpecies is called when production species is entered.
func (s *BaseSpeciesListener) EnterSpecies(ctx *SpeciesContext) {}

// ExitSpecies is called when production species is exited.
func (s *BaseSpeciesListener) ExitSpecies(ctx *SpeciesContext) {}
