// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672631742428/Term.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Term

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTermListener is a complete listener for a parse tree produced by TermParser.
type BaseTermListener struct{}

var _ TermListener = &BaseTermListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTermListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTermListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTermListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTermListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTermListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTermListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseTermListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseTermListener) ExitTerm(ctx *TermContext) {}
