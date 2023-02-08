// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675833204170/AddedSugars.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AddedSugars

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAddedSugarsListener is a complete listener for a parse tree produced by AddedSugarsParser.
type BaseAddedSugarsListener struct{}

var _ AddedSugarsListener = &BaseAddedSugarsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAddedSugarsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAddedSugarsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAddedSugarsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAddedSugarsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAddedSugarsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAddedSugarsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAddedsugars is called when production addedsugars is entered.
func (s *BaseAddedSugarsListener) EnterAddedsugars(ctx *AddedsugarsContext) {}

// ExitAddedsugars is called when production addedsugars is exited.
func (s *BaseAddedSugarsListener) ExitAddedsugars(ctx *AddedsugarsContext) {}
