// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867958192/Plastic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Plastic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePlasticListener is a complete listener for a parse tree produced by PlasticParser.
type BasePlasticListener struct{}

var _ PlasticListener = &BasePlasticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlasticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlasticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlasticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlasticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePlasticListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePlasticListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPlastic is called when production plastic is entered.
func (s *BasePlasticListener) EnterPlastic(ctx *PlasticContext) {}

// ExitPlastic is called when production plastic is exited.
func (s *BasePlasticListener) ExitPlastic(ctx *PlasticContext) {}
