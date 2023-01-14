// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673667498609/Domestic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Domestic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDomesticListener is a complete listener for a parse tree produced by DomesticParser.
type BaseDomesticListener struct{}

var _ DomesticListener = &BaseDomesticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDomesticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDomesticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDomesticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDomesticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDomesticListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDomesticListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDomestic is called when production domestic is entered.
func (s *BaseDomesticListener) EnterDomestic(ctx *DomesticContext) {}

// ExitDomestic is called when production domestic is exited.
func (s *BaseDomesticListener) ExitDomestic(ctx *DomesticContext) {}
