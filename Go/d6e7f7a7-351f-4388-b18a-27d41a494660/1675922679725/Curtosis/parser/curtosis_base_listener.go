// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675922679725/Curtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Curtosis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCurtosisListener is a complete listener for a parse tree produced by CurtosisParser.
type BaseCurtosisListener struct{}

var _ CurtosisListener = &BaseCurtosisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCurtosisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCurtosisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCurtosisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCurtosisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCurtosisListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCurtosisListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCurtosis is called when production curtosis is entered.
func (s *BaseCurtosisListener) EnterCurtosis(ctx *CurtosisContext) {}

// ExitCurtosis is called when production curtosis is exited.
func (s *BaseCurtosisListener) ExitCurtosis(ctx *CurtosisContext) {}
