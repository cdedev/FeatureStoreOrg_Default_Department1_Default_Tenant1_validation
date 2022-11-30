// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795008846/Biomass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Biomass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBiomassListener is a complete listener for a parse tree produced by BiomassParser.
type BaseBiomassListener struct{}

var _ BiomassListener = &BaseBiomassListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBiomassListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBiomassListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBiomassListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBiomassListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBiomassListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBiomassListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBiomass is called when production biomass is entered.
func (s *BaseBiomassListener) EnterBiomass(ctx *BiomassContext) {}

// ExitBiomass is called when production biomass is exited.
func (s *BaseBiomassListener) ExitBiomass(ctx *BiomassContext) {}
