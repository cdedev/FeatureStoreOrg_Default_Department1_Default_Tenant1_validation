// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672460775520/Structure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Structure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStructureListener is a complete listener for a parse tree produced by StructureParser.
type BaseStructureListener struct{}

var _ StructureListener = &BaseStructureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStructureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStructureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStructureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStructureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStructureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStructureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStructure is called when production structure is entered.
func (s *BaseStructureListener) EnterStructure(ctx *StructureContext) {}

// ExitStructure is called when production structure is exited.
func (s *BaseStructureListener) ExitStructure(ctx *StructureContext) {}
