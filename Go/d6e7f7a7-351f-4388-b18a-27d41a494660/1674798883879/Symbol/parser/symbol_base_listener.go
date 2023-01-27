// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674798883879/Symbol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Symbol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSymbolListener is a complete listener for a parse tree produced by SymbolParser.
type BaseSymbolListener struct{}

var _ SymbolListener = &BaseSymbolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSymbolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSymbolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSymbolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSymbolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSymbolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSymbolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseSymbolListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseSymbolListener) ExitSymbol(ctx *SymbolContext) {}
