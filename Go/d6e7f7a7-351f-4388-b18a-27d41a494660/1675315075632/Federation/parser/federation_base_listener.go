// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675315075632/Federation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Federation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFederationListener is a complete listener for a parse tree produced by FederationParser.
type BaseFederationListener struct{}

var _ FederationListener = &BaseFederationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFederationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFederationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFederationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFederationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFederationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFederationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFederation is called when production federation is entered.
func (s *BaseFederationListener) EnterFederation(ctx *FederationContext) {}

// ExitFederation is called when production federation is exited.
func (s *BaseFederationListener) ExitFederation(ctx *FederationContext) {}
