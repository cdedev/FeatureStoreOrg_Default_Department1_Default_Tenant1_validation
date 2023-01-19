// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674097827714/Registration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Registration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRegistrationListener is a complete listener for a parse tree produced by RegistrationParser.
type BaseRegistrationListener struct{}

var _ RegistrationListener = &BaseRegistrationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRegistrationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRegistrationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRegistrationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRegistrationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRegistrationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRegistrationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRegistration is called when production registration is entered.
func (s *BaseRegistrationListener) EnterRegistration(ctx *RegistrationContext) {}

// ExitRegistration is called when production registration is exited.
func (s *BaseRegistrationListener) ExitRegistration(ctx *RegistrationContext) {}
