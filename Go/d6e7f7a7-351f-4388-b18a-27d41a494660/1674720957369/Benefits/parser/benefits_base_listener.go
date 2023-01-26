// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674720957369/Benefits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Benefits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBenefitsListener is a complete listener for a parse tree produced by BenefitsParser.
type BaseBenefitsListener struct{}

var _ BenefitsListener = &BaseBenefitsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBenefitsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBenefitsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBenefitsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBenefitsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBenefitsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBenefitsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBenefits is called when production benefits is entered.
func (s *BaseBenefitsListener) EnterBenefits(ctx *BenefitsContext) {}

// ExitBenefits is called when production benefits is exited.
func (s *BaseBenefitsListener) ExitBenefits(ctx *BenefitsContext) {}
