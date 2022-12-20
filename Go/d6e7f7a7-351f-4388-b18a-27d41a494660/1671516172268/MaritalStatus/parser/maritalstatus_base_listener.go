// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516172268/MaritalStatus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaritalStatus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMaritalStatusListener is a complete listener for a parse tree produced by MaritalStatusParser.
type BaseMaritalStatusListener struct{}

var _ MaritalStatusListener = &BaseMaritalStatusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMaritalStatusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMaritalStatusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMaritalStatusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMaritalStatusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMaritalStatusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMaritalStatusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMaritalstatus is called when production maritalstatus is entered.
func (s *BaseMaritalStatusListener) EnterMaritalstatus(ctx *MaritalstatusContext) {}

// ExitMaritalstatus is called when production maritalstatus is exited.
func (s *BaseMaritalStatusListener) ExitMaritalstatus(ctx *MaritalstatusContext) {}
