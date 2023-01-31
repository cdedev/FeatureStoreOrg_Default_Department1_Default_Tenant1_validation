// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145203857/Words.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Words

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWordsListener is a complete listener for a parse tree produced by WordsParser.
type BaseWordsListener struct{}

var _ WordsListener = &BaseWordsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWordsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWordsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWordsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWordsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWordsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWordsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWords is called when production words is entered.
func (s *BaseWordsListener) EnterWords(ctx *WordsContext) {}

// ExitWords is called when production words is exited.
func (s *BaseWordsListener) ExitWords(ctx *WordsContext) {}
