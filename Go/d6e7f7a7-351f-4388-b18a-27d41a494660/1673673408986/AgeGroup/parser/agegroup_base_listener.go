// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673673408986/AgeGroup.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AgeGroup

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgeGroupListener is a complete listener for a parse tree produced by AgeGroupParser.
type BaseAgeGroupListener struct{}

var _ AgeGroupListener = &BaseAgeGroupListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgeGroupListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgeGroupListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgeGroupListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgeGroupListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAgeGroupListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAgeGroupListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAgegroup is called when production agegroup is entered.
func (s *BaseAgeGroupListener) EnterAgegroup(ctx *AgegroupContext) {}

// ExitAgegroup is called when production agegroup is exited.
func (s *BaseAgeGroupListener) ExitAgegroup(ctx *AgegroupContext) {}
