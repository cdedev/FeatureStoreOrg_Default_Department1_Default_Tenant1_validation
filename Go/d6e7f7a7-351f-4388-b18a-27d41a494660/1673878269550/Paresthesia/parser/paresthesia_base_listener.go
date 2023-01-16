// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878269550/Paresthesia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Paresthesia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseParesthesiaListener is a complete listener for a parse tree produced by ParesthesiaParser.
type BaseParesthesiaListener struct{}

var _ ParesthesiaListener = &BaseParesthesiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParesthesiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParesthesiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParesthesiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParesthesiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseParesthesiaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseParesthesiaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterParesthesia is called when production paresthesia is entered.
func (s *BaseParesthesiaListener) EnterParesthesia(ctx *ParesthesiaContext) {}

// ExitParesthesia is called when production paresthesia is exited.
func (s *BaseParesthesiaListener) ExitParesthesia(ctx *ParesthesiaContext) {}
