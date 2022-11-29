// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690110344/AddedSugar.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AddedSugar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAddedSugarListener is a complete listener for a parse tree produced by AddedSugarParser.
type BaseAddedSugarListener struct{}

var _ AddedSugarListener = &BaseAddedSugarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAddedSugarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAddedSugarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAddedSugarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAddedSugarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAddedSugarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAddedSugarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAddedsugar is called when production addedsugar is entered.
func (s *BaseAddedSugarListener) EnterAddedsugar(ctx *AddedsugarContext) {}

// ExitAddedsugar is called when production addedsugar is exited.
func (s *BaseAddedSugarListener) ExitAddedsugar(ctx *AddedsugarContext) {}
