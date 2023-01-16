// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868669341/Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClarityListener is a complete listener for a parse tree produced by ClarityParser.
type BaseClarityListener struct{}

var _ ClarityListener = &BaseClarityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClarityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClarityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClarityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClarityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseClarityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseClarityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterClarity is called when production clarity is entered.
func (s *BaseClarityListener) EnterClarity(ctx *ClarityContext) {}

// ExitClarity is called when production clarity is exited.
func (s *BaseClarityListener) ExitClarity(ctx *ClarityContext) {}
