// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447855819/Sex.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSexListener is a complete listener for a parse tree produced by SexParser.
type BaseSexListener struct{}

var _ SexListener = &BaseSexListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSexListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSexListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSexListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSexListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSexListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSexListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSex is called when production sex is entered.
func (s *BaseSexListener) EnterSex(ctx *SexContext) {}

// ExitSex is called when production sex is exited.
func (s *BaseSexListener) ExitSex(ctx *SexContext) {}
