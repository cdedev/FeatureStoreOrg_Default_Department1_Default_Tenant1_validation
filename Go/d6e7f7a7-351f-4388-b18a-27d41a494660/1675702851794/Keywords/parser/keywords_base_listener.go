// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675702851794/Keywords.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Keywords

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKeywordsListener is a complete listener for a parse tree produced by KeywordsParser.
type BaseKeywordsListener struct{}

var _ KeywordsListener = &BaseKeywordsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKeywordsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKeywordsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKeywordsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKeywordsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKeywordsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKeywordsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseKeywordsListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseKeywordsListener) ExitKeywords(ctx *KeywordsContext) {}
