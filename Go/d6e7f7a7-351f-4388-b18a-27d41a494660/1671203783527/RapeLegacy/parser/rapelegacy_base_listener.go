// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203783527/RapeLegacy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RapeLegacy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRapeLegacyListener is a complete listener for a parse tree produced by RapeLegacyParser.
type BaseRapeLegacyListener struct{}

var _ RapeLegacyListener = &BaseRapeLegacyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRapeLegacyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRapeLegacyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRapeLegacyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRapeLegacyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRapeLegacyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRapeLegacyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRapelegacy is called when production rapelegacy is entered.
func (s *BaseRapeLegacyListener) EnterRapelegacy(ctx *RapelegacyContext) {}

// ExitRapelegacy is called when production rapelegacy is exited.
func (s *BaseRapeLegacyListener) ExitRapelegacy(ctx *RapelegacyContext) {}
