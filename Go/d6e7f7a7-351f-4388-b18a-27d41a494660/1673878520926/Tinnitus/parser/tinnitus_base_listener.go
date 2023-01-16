// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878520926/Tinnitus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinnitus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTinnitusListener is a complete listener for a parse tree produced by TinnitusParser.
type BaseTinnitusListener struct{}

var _ TinnitusListener = &BaseTinnitusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTinnitusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTinnitusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTinnitusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTinnitusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTinnitusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTinnitusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTinnitus is called when production tinnitus is entered.
func (s *BaseTinnitusListener) EnterTinnitus(ctx *TinnitusContext) {}

// ExitTinnitus is called when production tinnitus is exited.
func (s *BaseTinnitusListener) ExitTinnitus(ctx *TinnitusContext) {}
