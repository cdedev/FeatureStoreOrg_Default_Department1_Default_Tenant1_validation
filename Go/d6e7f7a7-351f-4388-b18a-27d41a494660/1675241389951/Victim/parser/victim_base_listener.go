// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241389951/Victim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Victim

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVictimListener is a complete listener for a parse tree produced by VictimParser.
type BaseVictimListener struct{}

var _ VictimListener = &BaseVictimListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVictimListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVictimListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVictimListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVictimListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVictimListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVictimListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVictim is called when production victim is entered.
func (s *BaseVictimListener) EnterVictim(ctx *VictimContext) {}

// ExitVictim is called when production victim is exited.
func (s *BaseVictimListener) ExitVictim(ctx *VictimContext) {}
