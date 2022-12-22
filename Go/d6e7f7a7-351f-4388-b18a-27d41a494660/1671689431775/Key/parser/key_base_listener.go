// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689431775/Key.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Key

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKeyListener is a complete listener for a parse tree produced by KeyParser.
type BaseKeyListener struct{}

var _ KeyListener = &BaseKeyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKeyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKeyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKeyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKeyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKeyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKeyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKey is called when production key is entered.
func (s *BaseKeyListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseKeyListener) ExitKey(ctx *KeyContext) {}
