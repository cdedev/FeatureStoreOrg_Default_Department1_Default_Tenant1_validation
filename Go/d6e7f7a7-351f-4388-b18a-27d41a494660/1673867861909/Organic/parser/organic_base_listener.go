// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867861909/Organic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Organic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOrganicListener is a complete listener for a parse tree produced by OrganicParser.
type BaseOrganicListener struct{}

var _ OrganicListener = &BaseOrganicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOrganicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOrganicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOrganicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOrganicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOrganicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOrganicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrganic is called when production organic is entered.
func (s *BaseOrganicListener) EnterOrganic(ctx *OrganicContext) {}

// ExitOrganic is called when production organic is exited.
func (s *BaseOrganicListener) ExitOrganic(ctx *OrganicContext) {}
