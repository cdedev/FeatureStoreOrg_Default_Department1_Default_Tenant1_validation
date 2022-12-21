// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603542216/Aluminium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aluminium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAluminiumListener is a complete listener for a parse tree produced by AluminiumParser.
type BaseAluminiumListener struct{}

var _ AluminiumListener = &BaseAluminiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAluminiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAluminiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAluminiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAluminiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAluminiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAluminiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAluminium is called when production aluminium is entered.
func (s *BaseAluminiumListener) EnterAluminium(ctx *AluminiumContext) {}

// ExitAluminium is called when production aluminium is exited.
func (s *BaseAluminiumListener) ExitAluminium(ctx *AluminiumContext) {}
