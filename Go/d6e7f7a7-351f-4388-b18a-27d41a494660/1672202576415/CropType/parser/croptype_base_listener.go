// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202576415/CropType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CropType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCropTypeListener is a complete listener for a parse tree produced by CropTypeParser.
type BaseCropTypeListener struct{}

var _ CropTypeListener = &BaseCropTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCropTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCropTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCropTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCropTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCropTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCropTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCroptype is called when production croptype is entered.
func (s *BaseCropTypeListener) EnterCroptype(ctx *CroptypeContext) {}

// ExitCroptype is called when production croptype is exited.
func (s *BaseCropTypeListener) ExitCroptype(ctx *CroptypeContext) {}
