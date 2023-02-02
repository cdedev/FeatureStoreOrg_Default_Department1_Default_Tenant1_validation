// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675319648133/Code.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCodeListener is a complete listener for a parse tree produced by CodeParser.
type BaseCodeListener struct{}

var _ CodeListener = &BaseCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCode is called when production code is entered.
func (s *BaseCodeListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseCodeListener) ExitCode(ctx *CodeContext) {}
