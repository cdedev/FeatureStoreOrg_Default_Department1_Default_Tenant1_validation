// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669981752435/Skills.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Skills

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSkillsListener is a complete listener for a parse tree produced by SkillsParser.
type BaseSkillsListener struct{}

var _ SkillsListener = &BaseSkillsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSkillsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSkillsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSkillsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSkillsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSkillsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSkillsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSkills is called when production skills is entered.
func (s *BaseSkillsListener) EnterSkills(ctx *SkillsContext) {}

// ExitSkills is called when production skills is exited.
func (s *BaseSkillsListener) ExitSkills(ctx *SkillsContext) {}
