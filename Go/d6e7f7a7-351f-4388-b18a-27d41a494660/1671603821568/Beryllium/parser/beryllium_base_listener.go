// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603821568/Beryllium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Beryllium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBerylliumListener is a complete listener for a parse tree produced by BerylliumParser.
type BaseBerylliumListener struct{}

var _ BerylliumListener = &BaseBerylliumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBerylliumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBerylliumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBerylliumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBerylliumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBerylliumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBerylliumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBeryllium is called when production beryllium is entered.
func (s *BaseBerylliumListener) EnterBeryllium(ctx *BerylliumContext) {}

// ExitBeryllium is called when production beryllium is exited.
func (s *BaseBerylliumListener) ExitBeryllium(ctx *BerylliumContext) {}
