// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675757257083/Stage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStageListener is a complete listener for a parse tree produced by StageParser.
type BaseStageListener struct{}

var _ StageListener = &BaseStageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStage is called when production stage is entered.
func (s *BaseStageListener) EnterStage(ctx *StageContext) {}

// ExitStage is called when production stage is exited.
func (s *BaseStageListener) ExitStage(ctx *StageContext) {}
