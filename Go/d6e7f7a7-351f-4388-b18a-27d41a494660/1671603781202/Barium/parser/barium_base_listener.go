// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603781202/Barium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBariumListener is a complete listener for a parse tree produced by BariumParser.
type BaseBariumListener struct{}

var _ BariumListener = &BaseBariumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBariumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBariumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBariumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBariumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBariumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBariumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBarium is called when production barium is entered.
func (s *BaseBariumListener) EnterBarium(ctx *BariumContext) {}

// ExitBarium is called when production barium is exited.
func (s *BaseBariumListener) ExitBarium(ctx *BariumContext) {}
