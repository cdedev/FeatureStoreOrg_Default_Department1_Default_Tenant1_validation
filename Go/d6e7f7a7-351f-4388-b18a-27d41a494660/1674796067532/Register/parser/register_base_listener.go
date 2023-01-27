// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674796067532/Register.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Register

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRegisterListener is a complete listener for a parse tree produced by RegisterParser.
type BaseRegisterListener struct{}

var _ RegisterListener = &BaseRegisterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRegisterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRegisterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRegisterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRegisterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRegisterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRegisterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseRegisterListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseRegisterListener) ExitRegister(ctx *RegisterContext) {}
