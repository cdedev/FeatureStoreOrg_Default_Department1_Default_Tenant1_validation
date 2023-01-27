// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789940691/Cite.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cite

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCiteListener is a complete listener for a parse tree produced by CiteParser.
type BaseCiteListener struct{}

var _ CiteListener = &BaseCiteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCiteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCiteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCiteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCiteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCiteListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCiteListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCite is called when production cite is entered.
func (s *BaseCiteListener) EnterCite(ctx *CiteContext) {}

// ExitCite is called when production cite is exited.
func (s *BaseCiteListener) ExitCite(ctx *CiteContext) {}
