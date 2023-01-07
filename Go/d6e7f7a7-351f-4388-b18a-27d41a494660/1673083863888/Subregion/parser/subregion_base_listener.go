// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083863888/Subregion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Subregion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSubregionListener is a complete listener for a parse tree produced by SubregionParser.
type BaseSubregionListener struct{}

var _ SubregionListener = &BaseSubregionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSubregionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSubregionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSubregionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSubregionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSubregionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSubregionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubregion is called when production subregion is entered.
func (s *BaseSubregionListener) EnterSubregion(ctx *SubregionContext) {}

// ExitSubregion is called when production subregion is exited.
func (s *BaseSubregionListener) ExitSubregion(ctx *SubregionContext) {}
