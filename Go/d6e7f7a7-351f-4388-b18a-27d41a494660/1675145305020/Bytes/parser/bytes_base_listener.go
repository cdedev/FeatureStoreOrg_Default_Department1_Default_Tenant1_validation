// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145305020/Bytes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bytes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBytesListener is a complete listener for a parse tree produced by BytesParser.
type BaseBytesListener struct{}

var _ BytesListener = &BaseBytesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBytesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBytesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBytesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBytesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBytesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBytesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBytes is called when production bytes is entered.
func (s *BaseBytesListener) EnterBytes(ctx *BytesContext) {}

// ExitBytes is called when production bytes is exited.
func (s *BaseBytesListener) ExitBytes(ctx *BytesContext) {}
