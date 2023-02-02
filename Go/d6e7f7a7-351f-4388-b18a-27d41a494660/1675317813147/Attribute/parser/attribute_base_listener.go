// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675317813147/Attribute.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Attribute

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAttributeListener is a complete listener for a parse tree produced by AttributeParser.
type BaseAttributeListener struct{}

var _ AttributeListener = &BaseAttributeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAttributeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAttributeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAttributeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAttributeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAttributeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAttributeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseAttributeListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseAttributeListener) ExitAttribute(ctx *AttributeContext) {}
