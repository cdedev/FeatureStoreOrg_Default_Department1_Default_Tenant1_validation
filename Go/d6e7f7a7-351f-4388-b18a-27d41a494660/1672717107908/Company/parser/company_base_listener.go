// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717107908/Company.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Company

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCompanyListener is a complete listener for a parse tree produced by CompanyParser.
type BaseCompanyListener struct{}

var _ CompanyListener = &BaseCompanyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCompanyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCompanyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCompanyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCompanyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCompanyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCompanyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCompany is called when production company is entered.
func (s *BaseCompanyListener) EnterCompany(ctx *CompanyContext) {}

// ExitCompany is called when production company is exited.
func (s *BaseCompanyListener) ExitCompany(ctx *CompanyContext) {}
