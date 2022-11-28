// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626974586/Gender.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gender

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGenderListener is a complete listener for a parse tree produced by GenderParser.
type BaseGenderListener struct{}

var _ GenderListener = &BaseGenderListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGenderListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGenderListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGenderListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGenderListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGenderListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGenderListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGender is called when production gender is entered.
func (s *BaseGenderListener) EnterGender(ctx *GenderContext) {}

// ExitGender is called when production gender is exited.
func (s *BaseGenderListener) ExitGender(ctx *GenderContext) {}
