// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675932748157/CompanyType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCompanyTypeListener is a complete listener for a parse tree produced by CompanyTypeParser.
type BaseCompanyTypeListener struct{}

var _ CompanyTypeListener = &BaseCompanyTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCompanyTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCompanyTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCompanyTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCompanyTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCompanyTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCompanyTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCompanytype is called when production companytype is entered.
func (s *BaseCompanyTypeListener) EnterCompanytype(ctx *CompanytypeContext) {}

// ExitCompanytype is called when production companytype is exited.
func (s *BaseCompanyTypeListener) ExitCompanytype(ctx *CompanytypeContext) {}
