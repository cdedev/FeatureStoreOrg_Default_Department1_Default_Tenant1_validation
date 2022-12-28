// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672200271774/Default.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Default

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDefaultListener is a complete listener for a parse tree produced by DefaultParser.
type BaseDefaultListener struct{}

var _ DefaultListener = &BaseDefaultListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefaultListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefaultListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefaultListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefaultListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDefaultListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDefaultListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseDefaultListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseDefaultListener) ExitDefault(ctx *DefaultContext) {}
