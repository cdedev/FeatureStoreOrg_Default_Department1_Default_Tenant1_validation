// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201967666/Utterance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Utterance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUtteranceListener is a complete listener for a parse tree produced by UtteranceParser.
type BaseUtteranceListener struct{}

var _ UtteranceListener = &BaseUtteranceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUtteranceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUtteranceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUtteranceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUtteranceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUtteranceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUtteranceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUtterance is called when production utterance is entered.
func (s *BaseUtteranceListener) EnterUtterance(ctx *UtteranceContext) {}

// ExitUtterance is called when production utterance is exited.
func (s *BaseUtteranceListener) ExitUtterance(ctx *UtteranceContext) {}
