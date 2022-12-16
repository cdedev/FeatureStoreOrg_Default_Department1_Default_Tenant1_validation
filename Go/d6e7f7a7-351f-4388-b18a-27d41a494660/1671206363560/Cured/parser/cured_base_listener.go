// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206363560/Cured.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cured

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCuredListener is a complete listener for a parse tree produced by CuredParser.
type BaseCuredListener struct{}

var _ CuredListener = &BaseCuredListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCuredListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCuredListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCuredListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCuredListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCuredListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCuredListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCured is called when production cured is entered.
func (s *BaseCuredListener) EnterCured(ctx *CuredContext) {}

// ExitCured is called when production cured is exited.
func (s *BaseCuredListener) ExitCured(ctx *CuredContext) {}
