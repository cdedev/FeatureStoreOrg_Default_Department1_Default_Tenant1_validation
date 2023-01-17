// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925422241/Eggs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eggs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEggsListener is a complete listener for a parse tree produced by EggsParser.
type BaseEggsListener struct{}

var _ EggsListener = &BaseEggsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEggsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEggsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEggsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEggsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEggsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEggsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEggs is called when production eggs is entered.
func (s *BaseEggsListener) EnterEggs(ctx *EggsContext) {}

// ExitEggs is called when production eggs is exited.
func (s *BaseEggsListener) ExitEggs(ctx *EggsContext) {}
