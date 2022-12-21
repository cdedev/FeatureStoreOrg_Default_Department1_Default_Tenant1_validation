// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601050396/Agency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgencyListener is a complete listener for a parse tree produced by AgencyParser.
type BaseAgencyListener struct{}

var _ AgencyListener = &BaseAgencyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgencyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgencyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgencyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgencyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAgencyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAgencyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAgency is called when production agency is entered.
func (s *BaseAgencyListener) EnterAgency(ctx *AgencyContext) {}

// ExitAgency is called when production agency is exited.
func (s *BaseAgencyListener) ExitAgency(ctx *AgencyContext) {}
