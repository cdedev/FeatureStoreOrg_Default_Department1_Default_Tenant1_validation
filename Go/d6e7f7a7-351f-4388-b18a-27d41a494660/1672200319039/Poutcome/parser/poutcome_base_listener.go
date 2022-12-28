// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672200319039/Poutcome.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Poutcome

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePoutcomeListener is a complete listener for a parse tree produced by PoutcomeParser.
type BasePoutcomeListener struct{}

var _ PoutcomeListener = &BasePoutcomeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePoutcomeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePoutcomeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePoutcomeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePoutcomeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePoutcomeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePoutcomeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPoutcome is called when production poutcome is entered.
func (s *BasePoutcomeListener) EnterPoutcome(ctx *PoutcomeContext) {}

// ExitPoutcome is called when production poutcome is exited.
func (s *BasePoutcomeListener) ExitPoutcome(ctx *PoutcomeContext) {}
