// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528680321/Residence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Residence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseResidenceListener is a complete listener for a parse tree produced by ResidenceParser.
type BaseResidenceListener struct{}

var _ ResidenceListener = &BaseResidenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseResidenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseResidenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseResidenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseResidenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseResidenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseResidenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterResidence is called when production residence is entered.
func (s *BaseResidenceListener) EnterResidence(ctx *ResidenceContext) {}

// ExitResidence is called when production residence is exited.
func (s *BaseResidenceListener) ExitResidence(ctx *ResidenceContext) {}
