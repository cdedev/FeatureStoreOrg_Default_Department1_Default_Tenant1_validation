// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676464347040/Mortality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMortalityListener is a complete listener for a parse tree produced by MortalityParser.
type BaseMortalityListener struct{}

var _ MortalityListener = &BaseMortalityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMortalityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMortalityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMortalityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMortalityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMortalityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMortalityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMortality is called when production mortality is entered.
func (s *BaseMortalityListener) EnterMortality(ctx *MortalityContext) {}

// ExitMortality is called when production mortality is exited.
func (s *BaseMortalityListener) ExitMortality(ctx *MortalityContext) {}
