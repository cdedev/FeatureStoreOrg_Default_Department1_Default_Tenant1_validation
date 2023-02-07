// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742249049/CompanyName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyName

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCompanyNameListener is a complete listener for a parse tree produced by CompanyNameParser.
type BaseCompanyNameListener struct{}

var _ CompanyNameListener = &BaseCompanyNameListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCompanyNameListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCompanyNameListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCompanyNameListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCompanyNameListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCompanyNameListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCompanyNameListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCompanyname is called when production companyname is entered.
func (s *BaseCompanyNameListener) EnterCompanyname(ctx *CompanynameContext) {}

// ExitCompanyname is called when production companyname is exited.
func (s *BaseCompanyNameListener) ExitCompanyname(ctx *CompanynameContext) {}
