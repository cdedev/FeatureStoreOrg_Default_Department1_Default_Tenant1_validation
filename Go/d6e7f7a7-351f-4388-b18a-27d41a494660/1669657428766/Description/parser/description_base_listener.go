// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657428766/Description.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Description

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDescriptionListener is a complete listener for a parse tree produced by DescriptionParser.
type BaseDescriptionListener struct{}

var _ DescriptionListener = &BaseDescriptionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDescriptionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDescriptionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDescriptionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDescriptionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDescriptionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDescriptionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseDescriptionListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseDescriptionListener) ExitDescription(ctx *DescriptionContext) {}
