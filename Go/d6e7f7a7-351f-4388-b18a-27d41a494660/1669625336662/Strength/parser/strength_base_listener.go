// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669625336662/Strength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Strength

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStrengthListener is a complete listener for a parse tree produced by StrengthParser.
type BaseStrengthListener struct{}

var _ StrengthListener = &BaseStrengthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStrengthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStrengthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStrengthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStrengthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStrengthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStrengthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStrength is called when production strength is entered.
func (s *BaseStrengthListener) EnterStrength(ctx *StrengthContext) {}

// ExitStrength is called when production strength is exited.
func (s *BaseStrengthListener) ExitStrength(ctx *StrengthContext) {}
