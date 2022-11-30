// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669779000828/MultipleLines.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MultipleLines

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMultipleLinesListener is a complete listener for a parse tree produced by MultipleLinesParser.
type BaseMultipleLinesListener struct{}

var _ MultipleLinesListener = &BaseMultipleLinesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMultipleLinesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMultipleLinesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMultipleLinesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMultipleLinesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMultipleLinesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMultipleLinesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplelines is called when production multiplelines is entered.
func (s *BaseMultipleLinesListener) EnterMultiplelines(ctx *MultiplelinesContext) {}

// ExitMultiplelines is called when production multiplelines is exited.
func (s *BaseMultipleLinesListener) ExitMultiplelines(ctx *MultiplelinesContext) {}
