// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700185356/Meditation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Meditation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMeditationListener is a complete listener for a parse tree produced by MeditationParser.
type BaseMeditationListener struct{}

var _ MeditationListener = &BaseMeditationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMeditationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMeditationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMeditationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMeditationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMeditationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMeditationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMeditation is called when production meditation is entered.
func (s *BaseMeditationListener) EnterMeditation(ctx *MeditationContext) {}

// ExitMeditation is called when production meditation is exited.
func (s *BaseMeditationListener) ExitMeditation(ctx *MeditationContext) {}
