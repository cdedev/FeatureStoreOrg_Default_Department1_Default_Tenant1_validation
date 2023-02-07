// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742695319/Dividend.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dividend

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDividendListener is a complete listener for a parse tree produced by DividendParser.
type BaseDividendListener struct{}

var _ DividendListener = &BaseDividendListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDividendListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDividendListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDividendListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDividendListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDividendListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDividendListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDividend is called when production dividend is entered.
func (s *BaseDividendListener) EnterDividend(ctx *DividendContext) {}

// ExitDividend is called when production dividend is exited.
func (s *BaseDividendListener) ExitDividend(ctx *DividendContext) {}
