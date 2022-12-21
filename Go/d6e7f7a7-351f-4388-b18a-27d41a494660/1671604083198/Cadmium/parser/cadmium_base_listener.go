// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604083198/Cadmium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cadmium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCadmiumListener is a complete listener for a parse tree produced by CadmiumParser.
type BaseCadmiumListener struct{}

var _ CadmiumListener = &BaseCadmiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCadmiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCadmiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCadmiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCadmiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCadmiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCadmiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCadmium is called when production cadmium is entered.
func (s *BaseCadmiumListener) EnterCadmium(ctx *CadmiumContext) {}

// ExitCadmium is called when production cadmium is exited.
func (s *BaseCadmiumListener) ExitCadmium(ctx *CadmiumContext) {}
