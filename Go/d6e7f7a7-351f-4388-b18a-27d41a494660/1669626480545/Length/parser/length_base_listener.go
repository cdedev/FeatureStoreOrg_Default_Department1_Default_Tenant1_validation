// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626480545/Length.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Length

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLengthListener is a complete listener for a parse tree produced by LengthParser.
type BaseLengthListener struct{}

var _ LengthListener = &BaseLengthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLengthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLengthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLengthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLengthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLengthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLengthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLength is called when production length is entered.
func (s *BaseLengthListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseLengthListener) ExitLength(ctx *LengthContext) {}
