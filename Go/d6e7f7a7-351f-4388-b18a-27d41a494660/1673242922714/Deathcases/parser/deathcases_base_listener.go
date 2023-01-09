// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673242922714/Deathcases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Deathcases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDeathcasesListener is a complete listener for a parse tree produced by DeathcasesParser.
type BaseDeathcasesListener struct{}

var _ DeathcasesListener = &BaseDeathcasesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDeathcasesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDeathcasesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDeathcasesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDeathcasesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDeathcasesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDeathcasesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDeathcases is called when production deathcases is entered.
func (s *BaseDeathcasesListener) EnterDeathcases(ctx *DeathcasesContext) {}

// ExitDeathcases is called when production deathcases is exited.
func (s *BaseDeathcasesListener) ExitDeathcases(ctx *DeathcasesContext) {}
