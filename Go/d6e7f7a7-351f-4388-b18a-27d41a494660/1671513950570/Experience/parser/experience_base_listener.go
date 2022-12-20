// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671513950570/Experience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Experience

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExperienceListener is a complete listener for a parse tree produced by ExperienceParser.
type BaseExperienceListener struct{}

var _ ExperienceListener = &BaseExperienceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExperienceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExperienceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExperienceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExperienceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExperienceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExperienceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExperience is called when production experience is entered.
func (s *BaseExperienceListener) EnterExperience(ctx *ExperienceContext) {}

// ExitExperience is called when production experience is exited.
func (s *BaseExperienceListener) ExitExperience(ctx *ExperienceContext) {}
