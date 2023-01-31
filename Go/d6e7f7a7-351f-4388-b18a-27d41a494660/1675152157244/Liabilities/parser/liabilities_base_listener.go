// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675152157244/Liabilities.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liabilities

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLiabilitiesListener is a complete listener for a parse tree produced by LiabilitiesParser.
type BaseLiabilitiesListener struct{}

var _ LiabilitiesListener = &BaseLiabilitiesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLiabilitiesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLiabilitiesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLiabilitiesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLiabilitiesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLiabilitiesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLiabilitiesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLiabilities is called when production liabilities is entered.
func (s *BaseLiabilitiesListener) EnterLiabilities(ctx *LiabilitiesContext) {}

// ExitLiabilities is called when production liabilities is exited.
func (s *BaseLiabilitiesListener) ExitLiabilities(ctx *LiabilitiesContext) {}
