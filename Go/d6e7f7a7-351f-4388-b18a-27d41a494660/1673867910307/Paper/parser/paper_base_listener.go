// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867910307/Paper.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Paper

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePaperListener is a complete listener for a parse tree produced by PaperParser.
type BasePaperListener struct{}

var _ PaperListener = &BasePaperListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePaperListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePaperListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePaperListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePaperListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePaperListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePaperListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPaper is called when production paper is entered.
func (s *BasePaperListener) EnterPaper(ctx *PaperContext) {}

// ExitPaper is called when production paper is exited.
func (s *BasePaperListener) ExitPaper(ctx *PaperContext) {}
