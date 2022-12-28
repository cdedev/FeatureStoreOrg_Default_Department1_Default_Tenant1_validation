// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202276107/Charisma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Charisma

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCharismaListener is a complete listener for a parse tree produced by CharismaParser.
type BaseCharismaListener struct{}

var _ CharismaListener = &BaseCharismaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCharismaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCharismaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCharismaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCharismaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCharismaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCharismaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCharisma is called when production charisma is entered.
func (s *BaseCharismaListener) EnterCharisma(ctx *CharismaContext) {}

// ExitCharisma is called when production charisma is exited.
func (s *BaseCharismaListener) ExitCharisma(ctx *CharismaContext) {}
