// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673854633204/Model.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Model

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseModelListener is a complete listener for a parse tree produced by ModelParser.
type BaseModelListener struct{}

var _ ModelListener = &BaseModelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseModelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseModelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseModelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseModelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseModelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseModelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterModel is called when production model is entered.
func (s *BaseModelListener) EnterModel(ctx *ModelContext) {}

// ExitModel is called when production model is exited.
func (s *BaseModelListener) ExitModel(ctx *ModelContext) {}
