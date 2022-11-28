// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654511765/Field.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Field

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFieldListener is a complete listener for a parse tree produced by FieldParser.
type BaseFieldListener struct{}

var _ FieldListener = &BaseFieldListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFieldListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFieldListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFieldListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFieldListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFieldListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFieldListener) ExitExpression(ctx *ExpressionContext) {}

// EnterField is called when production field is entered.
func (s *BaseFieldListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseFieldListener) ExitField(ctx *FieldContext) {}
