// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878467126/Sensory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sensory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSensoryListener is a complete listener for a parse tree produced by SensoryParser.
type BaseSensoryListener struct{}

var _ SensoryListener = &BaseSensoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSensoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSensoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSensoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSensoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSensoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSensoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSensory is called when production sensory is entered.
func (s *BaseSensoryListener) EnterSensory(ctx *SensoryContext) {}

// ExitSensory is called when production sensory is exited.
func (s *BaseSensoryListener) ExitSensory(ctx *SensoryContext) {}
