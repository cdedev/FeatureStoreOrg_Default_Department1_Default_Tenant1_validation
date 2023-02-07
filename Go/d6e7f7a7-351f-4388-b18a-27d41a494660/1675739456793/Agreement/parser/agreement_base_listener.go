// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675739456793/Agreement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgreementListener is a complete listener for a parse tree produced by AgreementParser.
type BaseAgreementListener struct{}

var _ AgreementListener = &BaseAgreementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgreementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgreementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgreementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgreementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAgreementListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAgreementListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAgreement is called when production agreement is entered.
func (s *BaseAgreementListener) EnterAgreement(ctx *AgreementContext) {}

// ExitAgreement is called when production agreement is exited.
func (s *BaseAgreementListener) ExitAgreement(ctx *AgreementContext) {}
