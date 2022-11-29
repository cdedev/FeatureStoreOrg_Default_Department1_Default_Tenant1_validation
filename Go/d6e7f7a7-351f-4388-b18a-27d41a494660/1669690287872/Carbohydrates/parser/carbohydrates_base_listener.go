// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690287872/Carbohydrates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbohydrates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCarbohydratesListener is a complete listener for a parse tree produced by CarbohydratesParser.
type BaseCarbohydratesListener struct{}

var _ CarbohydratesListener = &BaseCarbohydratesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCarbohydratesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCarbohydratesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCarbohydratesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCarbohydratesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCarbohydratesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCarbohydratesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCarbohydrates is called when production carbohydrates is entered.
func (s *BaseCarbohydratesListener) EnterCarbohydrates(ctx *CarbohydratesContext) {}

// ExitCarbohydrates is called when production carbohydrates is exited.
func (s *BaseCarbohydratesListener) ExitCarbohydrates(ctx *CarbohydratesContext) {}
