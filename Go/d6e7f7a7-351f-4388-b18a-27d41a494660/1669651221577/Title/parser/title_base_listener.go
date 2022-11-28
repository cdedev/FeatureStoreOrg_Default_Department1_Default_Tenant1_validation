// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669651221577/Title.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Title

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTitleListener is a complete listener for a parse tree produced by TitleParser.
type BaseTitleListener struct{}

var _ TitleListener = &BaseTitleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTitleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTitleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTitleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTitleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTitleListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTitleListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseTitleListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseTitleListener) ExitTitle(ctx *TitleContext) {}
