// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669977757479/Population.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Population

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePopulationListener is a complete listener for a parse tree produced by PopulationParser.
type BasePopulationListener struct{}

var _ PopulationListener = &BasePopulationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePopulationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePopulationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePopulationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePopulationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePopulationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePopulationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPopulation is called when production population is entered.
func (s *BasePopulationListener) EnterPopulation(ctx *PopulationContext) {}

// ExitPopulation is called when production population is exited.
func (s *BasePopulationListener) ExitPopulation(ctx *PopulationContext) {}
