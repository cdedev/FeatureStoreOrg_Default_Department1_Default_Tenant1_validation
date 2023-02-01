// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675240749630/Goalscorer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Goalscorer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGoalscorerListener is a complete listener for a parse tree produced by GoalscorerParser.
type BaseGoalscorerListener struct{}

var _ GoalscorerListener = &BaseGoalscorerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoalscorerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoalscorerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoalscorerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoalscorerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoalscorerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoalscorerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGoalscorer is called when production goalscorer is entered.
func (s *BaseGoalscorerListener) EnterGoalscorer(ctx *GoalscorerContext) {}

// ExitGoalscorer is called when production goalscorer is exited.
func (s *BaseGoalscorerListener) ExitGoalscorer(ctx *GoalscorerContext) {}
